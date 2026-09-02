/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"api-observer-collector/internal/daemon"
	"log"

	"github.com/spf13/cobra"
)

var background bool

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start API Observer Collector",
	Run: func(cmd *cobra.Command, args []string) {
		err := daemon.StartDaemon(background)
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	startCmd.Flags().BoolVarP(&background, "background", "b", true, "Run as a background process")
}
