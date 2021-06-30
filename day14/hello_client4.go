package main

import (
	"fmt"
	"go-learn/day14/hello"
	"log"
)

func main() {
	client, err := hello.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var send hello.String = hello.String{
		Value: "hello",
	}
	var reply hello.String
	err = client.Hello(&send, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}
