package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/yuriykis/bth-speaker-on/system"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	fmt.Println(asciBanner)

	upIntervalFlag := flag.Int(
		"up-interval",
		5,
		"Interval in seconds to check if device is up",
	)
	upInterval := time.Duration(*upIntervalFlag) * time.Second
	flag.Parse()
	var (
		dm  system.DeviceManager
		err error
	)
	switch system.SystemType(runtime.GOOS) {
	case system.MacSystemType:
		dm, err = system.NewMacDeviceManager()
	case system.LinuxSystemType:
		dm, err = system.NewLinuxDeviceManager()
	case system.WindowsSystemType:
		dm, err = system.NewWindowsDeviceManager()
	default:
		log.Fatal("Unknown system type")
	}

	dm = system.NewLoggingDeviceManagerMiddleware(dm)
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	dm.Start(ctx, upInterval)
}

var asciBanner = `
  ___  _____  _  _   ___                   _                ___   _  _ 
 | _ )|_   _|| || | / __| _ __  ___  __ _ | |__ ___  _ _   / _ \ | \| |
 | _ \  | |  | __ | \__ \| '_ \/ -_)/ _' || / // -_)| '_| | (_) || .' |
 |___/  |_|  |_||_| |___/| .__/\___|\__,_||_\_\\___||_|    \___/ |_|\_|
                         |_|                                           
`
