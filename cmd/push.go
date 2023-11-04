package cmd

import (
	"context"
	"grpc-stack/internal/stack"
	"grpc-stack/pb"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.Flags().Int32P("number", "n", -1, "Number (int32) to push")
	pushCmd.MarkFlagRequired("number")
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push an int32 on the stack",
	Long:  `Push an int32 on the stack and get the stack size in response`,
	Run: func(cmd *cobra.Command, args []string) {
		number, _ := cmd.Flags().GetInt32("number")
		grpcPush(number)
	},
}

func grpcPush(number int32) {
	cc, _ := stack.OpenConnection(hostname, port)
	defer cc.Close()

	client := pb.NewStackServiceClient(cc)
	request := &pb.PushServiceRequest{Number: number}

	resp, err := client.Push(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Receive response => stack size:%d", resp.Size)
}
