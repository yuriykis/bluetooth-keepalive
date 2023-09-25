package main

import (
	"log"
	"runtime"
	"time"
)

type SystemType string

const (
	UnknownSystemType SystemType = ""
	MacSystemType     SystemType = "darwin"
	LinuxSystemType   SystemType = "linux"
	WindowsSystemType SystemType = "windows"
)

func (s SystemType) osType(stype string) SystemType {
	return SystemType(stype)
}

type System interface {
	Devices() ([]*Device, error)
}

func main() {
	var (
		system System
		err    error
		stype  SystemType
	)

	switch stype.osType(runtime.GOOS) {
	case MacSystemType:
		system, err = NewMacSystem()
	case LinuxSystemType:
		system, err = NewLinuxSystem()
	case WindowsSystemType:
		system, err = NewWindowsSystem()
	default:
		log.Fatal("Unknown system type")
	}
	if err != nil {
		log.Fatal(err)
	}

	devices, err := system.Devices()
	if err != nil {
		log.Fatal(err)
	}
	upDevicesLoop(devices)
}

func upDevicesLoop(devices []*Device) {
	for {
		for _, d := range devices {
			d.Up()
		}
		time.Sleep(1 * time.Second)
	}
}
