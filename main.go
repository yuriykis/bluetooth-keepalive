package main

import (
	"fmt"
	"runtime"

	"github.com/yuriykis/bluetooth-keepalive/cmd"
	"github.com/yuriykis/bluetooth-keepalive/log"
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
______ _____ _   _     _   __                 ___  _ _           
| ___ \_   _| | | |   | | / /                / _ \| (_)          
| |_/ / | | | |_| |   | |/ /  ___  ___ _ __ / /_\ \ |___   _____ 
| ___ \ | | |  _  |   |    \ / _ \/ _ \ '_ \|  _  | | \ \ / / _ \
| |_/ / | | | | | |   | |\  \  __/  __/ |_) | | | | | |\ V /  __/
\____/  \_/ \_| |_/   \_| \_/\___|\___| .__/\_| |_/_|_| \_/ \___|
                                      | |                        
                                      |_|                           
`
