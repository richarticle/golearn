package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"net"
)

var port = flag.String("p", "12345", "port")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", "0.0.0.0:"+*port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		go HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			//fmt.Println(err)
			return
		}

		line = bytes.TrimRight(line, "\r\n")
		if bytes.Equal(line, []byte("exit")) {
			fmt.Fprintf(conn, "bye!\r\n")
			return
		}

		fmt.Fprintf(conn, "%s\r\n", line)
	}
}
