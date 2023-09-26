package main

import (
	"context"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

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

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	devices, err := dm.Devices()
	if err != nil {
		log.Fatal(err)
	}
	quit := make(chan struct{})
	go upDevicesLoop(devices, quit)
	<-ctx.Done()
	log.Info("Exiting main...")
	quit <- struct{}{}
	<-quit
	log.Info("Done exiting main...")
}

func upDevicesLoop(devices []device.Devicer, quit chan struct{}) {
	for {
		select {
		case <-quit:
			log.Info("Exiting upDevicesLoop...")
			quit <- struct{}{}
			return
		default:
			for _, d := range devices {
				d.Up()
			}
			time.Sleep(1 * time.Second)
		}
	}
}
