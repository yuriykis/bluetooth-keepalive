package main

import (
	"log"
	"runtime"
	"time"

	"github.com/yuriykis/bth-speaker-on/device"
	"github.com/yuriykis/bth-speaker-on/system"
)

func main() {
	var (
		dm    system.DeviceManager
		err   error
		sType system.SystemType
	)
	switch sType.OsType(runtime.GOOS) {
	case system.MacSystemType:
		dm, err = system.NewMacDeviceManager()
	case system.LinuxSystemType:
		dm, err = system.NewLinuxDeviceManager()
	case system.WindowsSystemType:
		dm, err = system.NewWindowsDeviceManager()
	default:
		log.Fatal("Unknown system type")
	}
	if err != nil {
		log.Fatal(err)
	}

	devices, err := dm.Devices()
	if err != nil {
		log.Fatal(err)
	}
	upDevicesLoop(devices)
}

func upDevicesLoop(devices []device.Devicer) {
	for {
		for _, d := range devices {
			d.Up()
		}
		time.Sleep(1 * time.Second)
	}
}
