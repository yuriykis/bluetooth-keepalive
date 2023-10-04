package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/yuriykis/bth-speaker-on/system"
)

const (
	logFile = "bth-speaker-on.log"
)

func main() {
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
		log = setupLogger()
	)
	// fix logging, should be initialized before DeviceManager is created
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

	dm = system.NewLoggingDeviceManagerMiddleware(dm, log)
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

func setupLogger() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetLevel(logrus.DebugLevel)
	l.Out = os.Stdout
	logFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		l.Out = logFile
	} else {
		l.Info("Failed to log to file, using default stderr")
	}
	return l
}

var asciBanner = `
  ___  _____  _  _   ___                   _                ___   _  _ 
 | _ )|_   _|| || | / __| _ __  ___  __ _ | |__ ___  _ _   / _ \ | \| |
 | _ \  | |  | __ | \__ \| '_ \/ -_)/ _' || / // -_)| '_| | (_) || .' |
 |___/  |_|  |_||_| |___/| .__/\___|\__,_||_\_\\___||_|    \___/ |_|\_|
                         |_|                                           
`
