package system

import (
	"bytes"
	"os/exec"

	"github.com/yuriykis/bth-speaker-on/device"
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
		devices = device.MakeMacDevices(string(out))
	}
	return devices, nil
}
