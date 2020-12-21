package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
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

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	checkError(err, true)
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, true)
	result, err := readFully(tcpConn)
	checkError(err, true)
	var t time.Time
	_, err = asn1.Unmarshal(result, &t)
	fmt.Fprintf(os.Stdout, "Time: %v\n", t.String())
	os.Exit(0)
}
