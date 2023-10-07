package cmd

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
	"github.com/yuriykis/bluetooth-keepalive/input"
	"github.com/yuriykis/bluetooth-keepalive/log"
	"github.com/yuriykis/bluetooth-keepalive/system"
)

const (
	DaemonLogFileName = "/tmp/bluetooth-keepalive-daemon.log"
	DaemonPidFileName = "/tmp/bluetooth-keepalive.pid"
	MainLogFileName   = "/tmp/bluetooth-keepalive.log"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"start"},
	Short:   "Start bluetooth-keepalive",
	Run: func(cmd *cobra.Command, args []string) {
		log.ClearLogFile()

		var (
			upIntervalFlag = cmd.Flags().Lookup("up-interval")
			upInterval     *input.UpInterval
			err            error
		)

		flag.Parse()
		if upIntervalFlag != nil {
			upInterval, err = input.ParseUpIntervalFlag(upIntervalFlag)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			upInterval = input.DefaultUpInterval()
			log.Info("No up-interval flag provided, using default value", upInterval.Value)
		}

		if err := runDaemon(upInterval.Duration()); err != nil {
			log.Fatal(err)
		}
	},
}

func run(upInterval time.Duration) {
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

func runDaemon(upInterval time.Duration) error {
	cntxt := &daemon.Context{
		PidFileName: DaemonPidFileName,
		PidFilePerm: 0644,
		LogFileName: DaemonLogFileName,
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        os.Args,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		return err
	}
	if d != nil {
		return nil
	}
	defer cntxt.Release()

	run(upInterval)

	return nil
}
