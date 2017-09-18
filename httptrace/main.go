package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"time"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   4 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}

	// Dump request
	fmt.Println("-------- Request --------")
	printBytes(httputil.DumpRequestOut(req, true))
	fmt.Println("-------- End of Request --------\n\n")

	// Set trace callback
	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			log.Printf("GetConn: %s\n", hostPort)
		},
		GotConn: func(info httptrace.GotConnInfo) {
			log.Printf("GotConn: %v\n", info)
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			log.Printf("DNSStart: %v\n", info)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			log.Printf("DNSDone: %+v\n", info)
		},
		ConnectStart: func(network, addr string) {
			log.Printf("ConnectStart: %s %s\n", network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			log.Printf("ConnectDone: %s %s %v\n", network, addr, err)
		},
		TLSHandshakeStart: func() {
			log.Printf("TLSHandshakeStart\n")
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			log.Printf("TLSHandshakeDone %v: \n", err)
			fmt.Printf("%s\n", getTLSInfo(state))
		},
		WroteHeaders: func() {
			log.Printf("WroteHeaders\n")
		},
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			log.Printf("WroteRequest: %v\n", info)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("\n\n")

	// Dump response
	fmt.Println("-------- Response --------")
	printBytes(httputil.DumpResponse(resp, true))
	fmt.Println("-------- End of Response --------\n\n")
}

func printBytes(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n", data)
	} else {
		log.Fatalf("%s\n", err)
	}
}

func getTLSInfo(state tls.ConnectionState) string {
	buf := &bytes.Buffer{}

	fmt.Fprintf(buf, "TLS version: ")
	switch state.Version {
	case tls.VersionSSL30:
		fmt.Fprintf(buf, "SSL3.0")
	case tls.VersionTLS10:
		fmt.Fprintf(buf, "TLS1.0")
	case tls.VersionTLS11:
		fmt.Fprintf(buf, "TLS1.1")
	case tls.VersionTLS12:
		fmt.Fprintf(buf, "TLS1.2")
	default:
		fmt.Fprintf(buf, "%x", state.Version)
	}
	fmt.Fprintln(buf)

	fmt.Fprintf(buf, "HandshakeComplete: %v\n", state.HandshakeComplete)

	fmt.Fprintf(buf, "DidResume: %v\n", state.DidResume)

	fmt.Fprintf(buf, "CipherSuite: ")
	switch state.CipherSuite {
	case tls.TLS_RSA_WITH_RC4_128_SHA:
		fmt.Fprintf(buf, "TLS_RSA_WITH_RC4_128_SHA")
	case tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA:
		fmt.Fprintf(buf, "TLS_RSA_WITH_3DES_EDE_CBC_SHA")
	case tls.TLS_RSA_WITH_AES_128_CBC_SHA:
		fmt.Fprintf(buf, "TLS_RSA_WITH_AES_128_CBC_SHA")
	case tls.TLS_RSA_WITH_AES_256_CBC_SHA:
		fmt.Fprintf(buf, "TLS_RSA_WITH_AES_256_CBC_SHA")
	case tls.TLS_RSA_WITH_AES_128_CBC_SHA256:
		fmt.Fprintf(buf, "TLS_RSA_WITH_AES_128_CBC_SHA256")
	case tls.TLS_RSA_WITH_AES_128_GCM_SHA256:
		fmt.Fprintf(buf, "TLS_RSA_WITH_AES_128_GCM_SHA256")
	case tls.TLS_RSA_WITH_AES_256_GCM_SHA384:
		fmt.Fprintf(buf, "TLS_RSA_WITH_AES_256_GCM_SHA384")
	case tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA:
		fmt.Fprintf(buf, "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:
		fmt.Fprintf(buf, "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:
		fmt.Fprintf(buf, "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA")
	case tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA:
		fmt.Fprintf(buf, "TLS_ECDHE_RSA_WITH_RC4_128_SHA")
	case tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA:
		fmt.Fprintf(buf, "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA")
	case tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:
		fmt.Fprintf(buf, "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA")
	case tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:
		fmt.Fprintf(buf, "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256:
		fmt.Fprintf(buf, "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256")
	case tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256:
		fmt.Fprintf(buf, "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256")
	case tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:
		fmt.Fprintf(buf, "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:
		fmt.Fprintf(buf, "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256")
	case tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:
		fmt.Fprintf(buf, "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384")
	case tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:
		fmt.Fprintf(buf, "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384")
	case tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305:
		fmt.Fprintf(buf, "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305")
	case tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305:
		fmt.Fprintf(buf, "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305")
	default:
		fmt.Fprintf(buf, "%x", state.CipherSuite)
	}
	fmt.Fprintln(buf)

	return buf.String()
}
