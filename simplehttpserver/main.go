package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port int
	dir  string
)

func init() {
	flag.IntVar(&port, "p", 8080, "listen port")
	flag.StringVar(&dir, "d", ".", "file server directory")
}

func main() {
	flag.Parse()
	http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir)))
}
