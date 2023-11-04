package stack

import (
	"context"
	"fmt"
	"grpc-stack/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Stack is a stack of numbers
var stack *Stack

func init() {
	stack = NewStack()
}

// Server is the GRPC server
type Server struct {
	pb.StackServiceServer
}

// Push pushes a number to the stack
func (s *Server) Push(ctx context.Context, req *pb.PushServiceRequest) (*pb.PushServiceReply, error) {
	log.Printf("Received: %v", req)

	stack.Push(req.Number)

	log.Print("stack : ", stack)

	return &pb.PushServiceReply{
		Size: stack.Size(),
	}, nil
}

// Pop pops a number from the stack
func (s *Server) Pop(ctx context.Context, req *pb.PopServiceRequest) (*pb.PopServiceReply, error) {
	log.Printf("Received: %v", req)

	number, empty := stack.Pop()

	log.Print("stack : ", stack)

	return &pb.PopServiceReply{
		Number: number,
		Empty:  empty,
	}, nil
}

// Serve starts the GRPC server
func Serve(port int) {
	log.Printf("GRPC server listening on :%d", port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterStackServiceServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
