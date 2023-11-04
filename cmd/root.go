/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// server hostname
var hostname string

// server port
var port int

var rootCmd = &cobra.Command{
	Use:   "stack",
	Short: "Sandbox project using GRPC",
	Long:  `Push/pop number in a stack using GRPC`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&hostname, "addr", "i", "localhost", "server hostname / address")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stack.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.AddCommand(pushCmd)
}
