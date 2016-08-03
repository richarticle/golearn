package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/richarticle/golearn/grpc/protocol"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())

	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := protocol.NewSummationClient(conn)
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go FireComputeSum(c, wg)
		go FireComputeSum(c, wg)
		time.Sleep(time.Second * 5)
	}

	wg.Wait()
}

func FireComputeSum(c protocol.SummationClient, wg *sync.WaitGroup) {
	req := &protocol.Request{}
	for i := 0; i < 4; i++ {
		req.Numbers = append(req.Numbers, rand.Int31n(10))
	}
	log.Printf("Try to compute sum of %v", req.Numbers)
	resp, err := c.ComputeSum(context.Background(), req)
	if err != nil {
		log.Printf("Failed to ComputeSum: %v", err)
	} else {
		log.Printf("Sum of %v is %d", req.Numbers, resp.Sum)
	}
	wg.Done()
}
