package main

import (
	"log"
	"net"

	"github.com/richarticle/golearn/grpc/protocol"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// server is used to implement protocol.SummationServer.
type server struct{}

// ComputeSum computes the summantion of numbers
func (s *server) ComputeSum(ctx context.Context, req *protocol.Request) (*protocol.Response, error) {
	log.Println("Numbers:", req.Numbers)
	var sum int32
	for _, v := range req.Numbers {
		sum += v
	}
	return &protocol.Response{Sum: sum}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protocol.RegisterSummationServer(s, &server{})
	s.Serve(lis)
}
