package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 { // go run test_ip.go 127.0.0.1
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String()) // 127.0.0.1
	}
	os.Exit(0)
}
