package main

import (
	"fmt"
	"strings"

	"github.com/andybrewer/mack"
)

type Device struct {
	Name string
	Type string
}

func (d *Device) String() string {
	return fmt.Sprintf("%s - %s", d.Name, d.Type)
}

func (d *Device) Up() error {
	mack.Say("Up")
	return nil
}

func makeDevices(output string) []*Device {
	var devices []*Device
	d := strings.FieldsFunc(output, func(r rune) bool {
		return r == '\n' || r == '\t'
	})
	for _, v := range d {
		dName := strings.Split(v, ":")[0]
		dType := strings.Split(v, ":")[1]
		dName = clearString(dName)
		dType = clearString(dType)
		devices = append(devices, &Device{
			Name: dName,
			Type: dType,
		})
	}
	return devices
}
