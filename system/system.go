package system

import "github.com/yuriykis/bth-speaker-on/device"

type SystemType string

const (
	UnknownSystemType SystemType = ""
	MacSystemType     SystemType = "darwin"
	LinuxSystemType   SystemType = "linux"
	WindowsSystemType SystemType = "windows"
)

func (s SystemType) OsType(stype string) SystemType {
	return SystemType(stype)
}

type DeviceManager interface {
	Devices() ([]device.Devicer, error)
}
