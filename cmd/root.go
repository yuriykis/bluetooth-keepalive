package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuriykis/bluetooth-keepalive/log"
)

var rootCmd = &cobra.Command{
	Use:     "bluetooth-keepalive",
	Aliases: []string{"bluetooth-keepalive"},
	Short:   "bluetooth-keepalive is a tool to keep bluetooth speaker on",
	Long:    "bluetooth-keepalive is a tool to keep bluetooth speaker on",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	startCmd.Flags().IntP(
		"up-interval",
		"u",
		5,
		"Interval in minutes to check if device is up",
	)
	rootCmd.AddCommand(startCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
