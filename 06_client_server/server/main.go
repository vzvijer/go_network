package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err)
		os.Exit(1)
	}
}

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	fmt.Println(tcpAddr.IP)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	buffer := [1024]byte{}
	for {
		conn, err := listener.Accept()
		checkError(err)
		n, err := conn.Read(buffer[0:])
		checkError(err)
		fmt.Println(string(buffer[0:]))
		n, err = conn.Write(buffer[0:n])
		checkError(err)
	}
}
