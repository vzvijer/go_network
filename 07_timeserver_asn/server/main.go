package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error, quitOnError bool) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		if quitOnError {
			os.Exit(1)
		}
	}
}

func main() {
	port := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	checkError(err, true)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, true)
	for {
		conn, err := listener.AcceptTCP()
		checkError(err, false)
		fmt.Fprintf(os.Stdout, "Handling request from %v\n", conn.RemoteAddr())

		now := time.Now()
		m, err := asn1.Marshal(now)
		checkError(err, false)
		_, err = conn.Write(m)
		checkError(err, false)
		conn.Close()
	}
}
