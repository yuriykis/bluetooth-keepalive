package main

const (
	awkScriptLinux = `{print $3}`
	execCmdLinux   = "hcitool"
	execArgsLinux  = "con"
)

type LinuxSystem struct {
	devices []*Device
}

func (s *LinuxSystem) Devices() ([]*Device, error) {
	return s.devices, nil
}

func NewLinuxSystem() (*LinuxSystem, error) {
	devices, err := discoverLinuxDevices()
	if err != nil {
		return nil, err
	}
	return &LinuxSystem{
		devices: devices,
	}, nil
}

func discoverLinuxDevices() ([]*Device, error) {
	// TODO: implement
	return nil, nil
}
