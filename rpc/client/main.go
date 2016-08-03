package main

import (
	"log"
	"math/rand"
	"net/rpc"
	"sync"
	"time"

	"github.com/richarticle/golearn/rpc/protocol"
)

func main() {
	rand.Seed(time.Now().Unix())

	client, err := rpc.Dial("tcp", ":12345")
	if err != nil {
		log.Fatalf("Dial error: %v", err)
	}
	defer client.Close()

	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go FireComputeSum(client, wg)
		go FireComputeSum(client, wg)
		time.Sleep(time.Second * 5)
	}

	wg.Wait()
}

func FireComputeSum(c *rpc.Client, wg *sync.WaitGroup) {
	req := &protocol.Request{}
	for i := 0; i < 4; i++ {
		req.Numbers = append(req.Numbers, rand.Int31n(10))
	}
	log.Printf("Try to compute sum of %v", req.Numbers)
	resp := new(protocol.Response)
	err := c.Call("SummationServer.ComputeSum", req, resp)
	if err != nil {
		log.Fatalf("ComputeSum error: %v", err)
	} else {
		log.Printf("Sum of %v is %d", req.Numbers, resp.Sum)
	}
	wg.Done()
}
