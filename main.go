package main

import (
	"context"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

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
	mainWg := &sync.WaitGroup{}
	mainWg.Add(1)
	go upDevicesLoop(ctx, devices, mainWg)
	mainWg.Wait()
	log.Info("Exiting main...")
	log.Info("Done exiting main...")
}

func upDevicesLoop(
	ctx context.Context,
	devices []device.Devicer,
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
