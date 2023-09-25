package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/andybrewer/mack"
)

const (
	Speaker = "Speaker"
)

type Device struct {
	Name string
	Type string
}

func (d *Device) MusicPlaying() (bool, error) {
	if d.Type != Speaker {
		return false, errors.New("device is not a speaker")
	}
	cmd1 := exec.Command("/usr/bin/pmset", "-g")
	cmd2 := exec.Command("grep", " sleep")
	output, err := cmd1.Output()
	if err != nil {
		return false, err
	}
	cmd2.Stdin = bytes.NewBuffer(output)
	output, err = cmd2.Output()
	if err != nil {
		return false, err
	}
	return strings.Contains(string(output), "coreaudiod"), nil
}

func (d *Device) String() string {
	return fmt.Sprintf("%s - %s", d.Name, d.Type)
}

func (d *Device) Up() error {
	ok, err := d.MusicPlaying()
	if err != nil {
		return err
	}
	if !ok {
		if err := d.Volume(0.2); err != nil {
			return err
		}
	}
	return nil
}

func (d *Device) Play() error {
	if d.Type != Speaker {
		return nil
	}
	if err := mack.Say("Up"); err != nil {
		return err
	}
	log.Println("Playing music")
	return nil
}

func (d *Device) Volume(v float32) error {
	cmd1 := exec.Command("osascript", "-e", "output volume of (get volume settings)")
	output, err := cmd1.Output()
	if err != nil {
		return err
	}
	currentVol := string(output)
	fmt.Println(currentVol)

	cmd2 := exec.Command("osascript", "-e", fmt.Sprintf("set volume output volume %f", v))
	log.Println("Setting volume to", v)
	err = cmd2.Run()
	if err != nil {
		return err
	}

	d.Play()

	cmd3 := exec.Command(
		"osascript",
		"-e",
		fmt.Sprintf("set volume output volume %s", currentVol),
	)
	log.Println("Setting volume back to", currentVol)
	err = cmd3.Run()
	if err != nil {
		return err
	}
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
