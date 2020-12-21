package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s domain ...", os.Args[0])
		os.Exit(1)
	}
	for _, domain := range os.Args[1:] {
		ip, err := net.ResolveIPAddr("ip", domain)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while resolving ip address of %s: %s\n", domain, err)
		} else {
			fmt.Fprintf(os.Stdout, "%s -> %s\n", domain, ip.String())
		}
	}
}
