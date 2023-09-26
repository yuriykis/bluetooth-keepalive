package system

import "github.com/yuriykis/bth-speaker-on/device"

const (
	awkScriptLinux = `{print $3}`
	execCmdLinux   = "hcitool"
	execArgsLinux  = "con"
)

type LinuxDeviceManager struct {
	devices []device.Devicer
}

func (s *LinuxDeviceManager) Devices() ([]device.Devicer, error) {
	return s.devices, nil
}

func NewLinuxDeviceManager() (*LinuxDeviceManager, error) {
	devices, err := discoverLinuxDevices()
	if err != nil {
		return nil, err
	}
	return &LinuxDeviceManager{
		devices: devices,
	}, nil
}

func discoverLinuxDevices() ([]device.Devicer, error) {
	// TODO: implement
	return nil, nil
}
