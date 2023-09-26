package device

import (
	"fmt"
	"strings"

	"github.com/yuriykis/bth-speaker-on/util"
)

type DeviceType string

const (
	UnknownDeviceType DeviceType = ""
	SpeakerType       DeviceType = "Speaker"
	KeybordType       DeviceType = "Keyboard"
)

type Devicer interface {
	Up() error
	String() string
}

func MakeDevices(output string) []Devicer {
	var devices []Devicer
	d := strings.FieldsFunc(output, func(r rune) bool {
		return r == '\n' || r == '\t'
	})
	for _, v := range d {
		dName := strings.Split(v, ":")[0]
		dType := strings.Split(v, ":")[1]
		dName = util.ClearString(dName)
		dType = util.ClearString(dType)
		fmt.Printf("Devicer: %s, Type: %s\n", dName, dType)

		s := NewSpeaker(dName)
		devices = append(devices, s)
	}
	return devices
}
