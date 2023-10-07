package system

import (
	"bytes"
	"context"
	"time"

	"os/exec"
	"strings"
	"sync"

	"github.com/yuriykis/bluetooth-keepalive/device"
	"github.com/yuriykis/bluetooth-keepalive/log"
	"github.com/yuriykis/bluetooth-keepalive/util"
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
	return &MacDeviceManager{}, nil
}

func (dm *MacDeviceManager) Start(ctx context.Context, d time.Duration) error {
	devices, err := discoverMacDevices()
	if err != nil {
		log.Fatal(err)
	}
	dm.devices = devices
	mainWg := &sync.WaitGroup{}
	mainWg.Add(1)
	go device.UpDevicesLoop(ctx, devices, mainWg, d)
	mainWg.Wait()

	return nil
}

func discoverMacDevices() ([]device.Devicer, error) {
	output, err := runCmd(execCmdMac, nil, execArgsMac)
	if err != nil {
		return nil, err
	}
	var devices []device.Devicer
	if len(awkScriptMac) > 0 {
		out, err := runCmd("awk", output, awkScriptMac)
		if err != nil {
			return nil, err
		}
		devices = MakeMacDevices(string(out))
	}
	log.WithFields(log.Fields{
		"devices": devices,
	}).Info("Discovered devices")

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

		log.WithFields(log.Fields{
			"system": "Mac OS",
			"device": dName,
			"type":   dType,
		}).Info("Discovered device")

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
			log.WithFields(log.Fields{
				"system": "Mac OS",
				"device": dName,
				"type":   dType,
			}).Info("Device not supported")
			continue
		}
		s = device.NewLoggingMiddleware(s)
		devices = append(devices, s)
	}
	return devices
}

func parseMacDevicesOutput(output string) []string {
	return strings.FieldsFunc(output, func(r rune) bool {
		return r == '\n' || r == '\t'
	})
}

func runCmd(cmd string, input []byte, args ...string) ([]byte, error) {
	command := exec.Command(cmd, args...)
	if len(input) > 0 {
		command.Stdin = bytes.NewBuffer(input)
	}
	output, err := command.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, nil
}
