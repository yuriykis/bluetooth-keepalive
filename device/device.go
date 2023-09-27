package device

import (
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/yuriykis/bth-speaker-on/util"
)

type Devicer interface {
	Up(wg *sync.WaitGroup) error
	String() string
}

func MakeMacDevices(output string) []Devicer {
	var devices []Devicer
	d := strings.FieldsFunc(output, func(r rune) bool {
		return r == '\n' || r == '\t'
	})
	for _, v := range d {
		dName := strings.Split(v, ":")[0]
		dType := strings.Split(v, ":")[1]
		dName = util.ClearString(dName)
		dType = util.ClearString(dType)
		log.Printf("System: Mac OS, Device: %s, Type: %s\n", dName, dType)

		s := NewSpeaker(dName)
		devices = append(devices, s)
	}
	return devices
}

func MakeLinuxDevices(output string) []Devicer {
	return nil
}

func MakeWindowsDevices(output string) []Devicer {
	return nil
}
