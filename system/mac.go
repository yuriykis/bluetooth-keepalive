package system

import (
	"bytes"
	"context"

	log "github.com/sirupsen/logrus"

	"os/exec"
	"strings"
	"sync"

	"github.com/yuriykis/bth-speaker-on/device"
	"github.com/yuriykis/bth-speaker-on/util"
)

const (
	awkScriptMac = `
			BEGIN {show=0; device=""; minorType=""; processing=0}
			/^ *Connected:/ {show=1; next}
			/^ *Not Connected:/ {exit}
			show && /Address:/ && device !~ /Bluetooth Controller/ {deviceName=device; processing=1}
			processing && /Minor Type:/ {minorType=$3; print deviceName, "-", minorType; deviceName=""; minorType=""; processing=0}
			{device=$0}
			`
	execCmdMac  = "system_profiler"
	execArgsMac = "SPBluetoothDataType"
)

type MacDeviceManager struct {
	devices []device.Devicer
}

func (s *MacDeviceManager) Devices() ([]device.Devicer, error) {
	return s.devices, nil
}

func NewMacDeviceManager() (*MacDeviceManager, error) {
	devices, err := discoverMacDevices()
	if err != nil {
		return nil, err
	}
	return &MacDeviceManager{
		devices: devices,
	}, nil
}

func (dm *MacDeviceManager) Start(ctx context.Context) error {
	devices, err := dm.Devices()
	if err != nil {
		log.Fatal(err)
	}

	mainWg := &sync.WaitGroup{}
	mainWg.Add(1)
	go device.UpDevicesLoop(ctx, devices, mainWg)
	mainWg.Wait()

	log.Info("Exiting main...")
	log.Info("Done exiting main...")
	return nil
}

func discoverMacDevices() ([]device.Devicer, error) {
	cmd := exec.Command(execCmdMac, execArgsMac)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var devices []device.Devicer
	if len(awkScriptMac) > 0 {
		cmd2 := exec.Command("awk", awkScriptMac)
		cmd2.Stdin = bytes.NewBuffer(output)
		out, err := cmd2.Output()
		if err != nil {
			return nil, err
		}
		devices = MakeMacDevices(string(out))
	}
	return devices, nil
}

func MakeMacDevices(output string) []device.Devicer {
	var devices []device.Devicer
	d := parseMacDevicesOutput(output)
	for _, v := range d {
		var (
			s     device.Devicer
			dName = util.ClearString(strings.Split(v, ":")[0])
			dType = util.ClearString(strings.Split(v, ":")[1])
		)

		log.Printf("System: Mac OS, Device: %s, Type: %s\n", dName, dType)

		switch device.DeviceType(dType) {
		case device.SpeakerDeviceType:
			s = device.NewSpeaker(dName)
		case device.KeybordDeviceType:
			// TODO: implement
			continue
		case device.MouseDeviceType:
			// TODO: implement
			continue
		default:
			log.Printf(
				"System: Mac OS, Device: %s, Type: %s, not supported\n",
				dName,
				dType,
			)
			continue
		}
		devices = append(devices, s)
	}
	return devices
}

func parseMacDevicesOutput(output string) []string {
	return strings.FieldsFunc(output, func(r rune) bool {
		return r == '\n' || r == '\t'
	})
}
