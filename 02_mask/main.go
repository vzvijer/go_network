package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr ones bits\n", os.Args[0])
		os.Exit(1)
	}

	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		fmt.Println("Invalid ip address.")
		os.Exit(1)
	}

	ones, err := strconv.Atoi(os.Args[2])
	if err != nil {

	}

	bits, err := strconv.Atoi(os.Args[3])
	if err != nil {

	}

	fmt.Println("Default mask: ", ip.DefaultMask().String())
	fmt.Println("Default network:", ip.Mask(ip.DefaultMask()))

	mask := net.CIDRMask(ones, bits)
	fmt.Println("Mask:", mask.String())
	fmt.Println("Network:", ip.Mask(mask).String())
}
