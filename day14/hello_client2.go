package main

import (
	"fmt"
	"go-learn/day14/rpc_pkg"
	"log"
)

func main() {
	client, err := rpc_pkg.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
