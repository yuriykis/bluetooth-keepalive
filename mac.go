package main

import (
	"bytes"
	"os/exec"
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

type MacSystem struct {
	devices []*Device
}

func (s *MacSystem) Devices() ([]*Device, error) {
	return s.devices, nil
}

func NewMacSystem() (*MacSystem, error) {
	devices, err := discoverMacDevices()
	if err != nil {
		return nil, err
	}
	return &MacSystem{
		devices: devices,
	}, nil
}

func discoverMacDevices() ([]*Device, error) {
	cmd := exec.Command(execCmdMac, execArgsMac)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var devices []*Device
	if len(awkScriptMac) > 0 {
		cmd2 := exec.Command("awk", awkScriptMac)
		cmd2.Stdin = bytes.NewBuffer(output)
		out, err := cmd2.Output()
		if err != nil {
			return nil, err
		}
		devices = makeDevices(string(out))
	}
	return devices, nil
}
