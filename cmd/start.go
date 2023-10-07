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
	"github.com/yuriykis/bth-speaker-on/input"
	"github.com/yuriykis/bth-speaker-on/log"
	"github.com/yuriykis/bth-speaker-on/system"
)

const (
	DaemonLogFileName = "/tmp/bth-speaker-on-daemon.log"
	DaemonPidFileName = "/tmp/bth-speaker-on.pid"
	MainLogFileName   = "/tmp/bth-speaker-on.log"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"start"},
	Short:   "Start bth-speaker-on",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.ClearLogFile()

		upIntervalFlag := cmd.Flags().Lookup("up-interval")
		if upIntervalFlag == nil {
			log.Fatal("up-interval flag is not set")
		}
		flag.Parse()
		upInterval, err := input.ParseUpIntervalFlag(upIntervalFlag)
		if err != nil {
			log.Fatal(err)
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
