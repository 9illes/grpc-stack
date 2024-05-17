package cmd

import (
	"context"
	"log"

	"github.com/9illes/grpc-stack/internal/stack"

	"github.com/9illes/grpc-stack/pb"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(popCmd)
}

var popCmd = &cobra.Command{
	Use:   "pop",
	Short: "pop an int32 from the stack",
	Long:  `pop an int32 from the stack`,
	Run: func(cmd *cobra.Command, args []string) {
		grpcPop()
	},
}

func grpcPop() {
	cc, _ := stack.OpenConnection(hostname, port)
	defer cc.Close()

	client := pb.NewStackServiceClient(cc)
	request := &pb.PopServiceRequest{}

	resp, err := client.Pop(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Empty {
		log.Print("Receive response => stack is empty")
	} else {
		log.Printf("Receive response => %d", resp.Number)
	}
}
