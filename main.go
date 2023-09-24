package main

import (
	"fmt"
	"log"
	"runtime"
)

type SystemType string

const (
	UnknownSystemType SystemType = ""
	MacSystemType     SystemType = "darwin"
	LinuxSystemType   SystemType = "linux"
	WindowsSystemType SystemType = "windows"
)

type System interface {
	Devices() ([]*Device, error)
}

func checkOSType() SystemType {
	switch runtime.GOOS {
	case "darwin":
		return MacSystemType
	case "linux":
		return LinuxSystemType
	case "windows":
		return WindowsSystemType
	default:
		return UnknownSystemType
	}
}

func main() {
	var (
		system System
		err    error
	)
	osSystemType := checkOSType()
	switch osSystemType {
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
	for _, d := range devices {
		fmt.Println(d)
	}
}
