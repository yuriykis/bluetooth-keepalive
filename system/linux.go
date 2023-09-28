package system

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/yuriykis/bth-speaker-on/device"
)

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

func (s *LinuxDeviceManager) Start(ctx context.Context) error {
	log.Println("LinuxDeviceManager.Start() is not implemented")
	return nil
}

func discoverLinuxDevices() ([]device.Devicer, error) {
	// TODO: implement
	return nil, nil
}

func MakeLinuxDevices(output string) []device.Devicer {
	// TODO: implement
	return nil
}
