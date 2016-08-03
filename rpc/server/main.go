package main

import (
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/richarticle/golearn/rpc/protocol"
)

type SummationServer struct{}

// ComputeSum computes the summantion of numbers
func (s *SummationServer) ComputeSum(req *protocol.Request, resp *protocol.Response) error {
	log.Println("Numbers:", req.Numbers)
	var sum int32
	for _, v := range req.Numbers {
		sum += v
	}
	resp.Sum = sum
	time.Sleep(time.Second)
	return nil
}

func main() {
	rpc.Register(new(SummationServer))
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	rpc.Accept(lis)
}
