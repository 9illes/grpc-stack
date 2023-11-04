package cmd

import (
	"grpc-stack/internal/stack"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start GRPC server",
	Long:  `Start GRPC Stack server on the specified port (default: 8080)`,
	Run: func(cmd *cobra.Command, args []string) {
		stack.Serve(port)
	},
}
