package system

import (
	"context"

	"github.com/yuriykis/bth-speaker-on/device"
)

const (
	awkScriptWindows = ``
	execCmdWindows   = ""
	execArgsWindows  = ""
)

type WindowsDeviceManager struct {
	devices []device.Devicer
}

func (s *WindowsDeviceManager) Devices() ([]device.Devicer, error) {
	return s.devices, nil
}

func NewWindowsDeviceManager() (*WindowsDeviceManager, error) {
	devices, err := discoverWindowsDevices()
	if err != nil {
		return nil, err
	}
	return &WindowsDeviceManager{
		devices: devices,
	}, nil
}

func (s *WindowsDeviceManager) Start(ctx context.Context) error {
	return nil
}

func discoverWindowsDevices() ([]device.Devicer, error) {
	// TODO: implement
	return nil, nil
}

func MakeWindowsDevices(output string) []device.Devicer {
	// TODO: implement
	return nil
}
