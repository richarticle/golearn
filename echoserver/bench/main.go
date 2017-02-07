package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var addr = flag.String("addr", "127.0.0.1:12345", "address ip:port")
var numQueries = flag.Int("n", 1, "number of messages for each connection")
var numConns = flag.Int("c", 1, "number of connections")
var msg = flag.String("m", "helloworld", "message to send")

func main() {
	flag.Parse()

	wg := new(sync.WaitGroup)
	wg.Add(*numConns)

	t1 := time.Now()

	// Start connections
	for i := 1; i <= *numConns; i++ {
		go StartConn(*addr, *numQueries, *msg, wg)
	}

	wg.Wait()

	secs := time.Since(t1).Seconds()

	performance := float64(*numQueries) * float64(*numConns) / secs

	fmt.Printf("Performance: %f messages/second\n", performance)

	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func StartConn(addr string, numQueries int, msg string, wg *sync.WaitGroup) {

	bytesMsg := []byte(msg)

	// Connect to server
	conn, err := net.Dial("tcp", addr)
	checkError(err)

	defer conn.Close()

	reader := bufio.NewReader(conn)

	for i := 1; i <= numQueries; i++ {
		fmt.Fprintf(conn, "%s\r\n", msg)
		resp, err := reader.ReadBytes('\n')
		checkError(err)

		resp = bytes.TrimRight(resp, "\r\n")
		if !bytes.Equal(resp, bytesMsg) {
			fmt.Printf("Error response %s\n", resp)
			break
		}

	}
	wg.Done()
}
