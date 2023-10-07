package main

import (
	"fmt"
	"runtime"

	"github.com/yuriykis/bth-speaker-on/cmd"
	"github.com/yuriykis/bth-speaker-on/log"
)

func printStartupInfo() {
	fmt.Println(asciBanner)
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)
	fmt.Printf("Version: %s\n", runtime.Version())
	fmt.Printf("Log file: %s\n", cmd.MainLogFileName)
	fmt.Printf("Daemon log file: %s\n", cmd.DaemonLogFileName)
	fmt.Printf("Daemon pid file: %s\n", cmd.DaemonPidFileName)
}

func main() {
	log.Println(asciBanner)
	printStartupInfo()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var asciBanner = `
  ___  _____  _  _   ___                   _                ___   _  _ 
 | _ )|_   _|| || | / __| _ __  ___  __ _ | |__ ___  _ _   / _ \ | \| |
 | _ \  | |  | __ | \__ \| '_ \/ -_)/ _' || / // -_)| '_| | (_) || .' |
 |___/  |_|  |_||_| |___/| .__/\___|\__,_||_\_\\___||_|    \___/ |_|\_|
                         |_|                                           
`
