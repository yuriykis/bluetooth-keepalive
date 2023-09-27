package device

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
)

type DeviceType string

const (
	UnknownDeviceType DeviceType = ""
	SpeakerDeviceType DeviceType = "Speaker"
	KeybordDeviceType DeviceType = "Keybord"
	MouseDeviceType   DeviceType = "Mouse"
)

type Devicer interface {
	Up(wg *sync.WaitGroup) error
	String() string
}

func UpDevicesLoop(
	ctx context.Context,
	devices []Devicer,
	mainWg *sync.WaitGroup,
) {
	for {
		select {
		case <-ctx.Done():
			log.Info("Exiting upDevicesLoop...")
			mainWg.Done()
			return
		default:
			wg := &sync.WaitGroup{}
			for _, d := range devices {
				log.Infof("Device: %s", d.String())
				wg.Add(1)
				go d.Up(wg)
			}
			log.Info("Waiting for devices to be up...")
			wg.Wait()
		}
	}
}
