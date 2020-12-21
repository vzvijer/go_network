package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-appr\n", os.Args[0])
		os.Exit(1)
	}
	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		fmt.Println("Invalid address.")
	} else {
		fmt.Println("The address is:", ip)
	}
	os.Exit(0)
}
