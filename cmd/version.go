package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GRPC-Stack",
	Long:  `Print the version number of GRPC-Stack`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GRPC stack v0.1.0")
	},
}
