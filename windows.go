package main

const (
	awkScriptWindows = ``
	execCmdWindows   = ""
	execArgsWindows  = ""
)

type WindowsSystem struct {
	devices []*Device
}

func (s *WindowsSystem) Devices() ([]*Device, error) {
	return s.devices, nil
}

func NewWindowsSystem() (*WindowsSystem, error) {
	devices, err := discoverWindowsDevices()
	if err != nil {
		return nil, err
	}
	return &WindowsSystem{
		devices: devices,
	}, nil
}

func discoverWindowsDevices() ([]*Device, error) {
	// TODO: implement
	return nil, nil
}
