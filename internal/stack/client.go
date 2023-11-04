package stack

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// OpenConnection opens a connection to a stack server. Don't forget to close the connection.
func OpenConnection(hostname string, port int) (*grpc.ClientConn, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(fmt.Sprintf("%s:%d", hostname, port), opts)
	if err != nil {
		log.Fatal(err)
	}

	return cc, nil
}
