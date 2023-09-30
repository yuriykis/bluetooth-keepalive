package system

import (
	"context"
	"time"

	"github.com/yuriykis/bth-speaker-on/device"
)

type SystemType string

const (
	UnknownSystemType SystemType = ""
	MacSystemType     SystemType = "darwin"
	LinuxSystemType   SystemType = "linux"
	WindowsSystemType SystemType = "windows"
)

type DeviceManager interface {
	Devices() ([]device.Devicer, error)
	Start(context.Context, time.Duration) error
}
