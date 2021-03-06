package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func doClientWork2(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer func(client *rpc.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	var reply string
	err := client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

func doClientWork3(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	helloCall := client.Go("HelloService.Hello", "hello", new(string), nil)

	// do some thing

	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}

	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)
	fmt.Println(args, *reply)
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error:", err)
			}

			clientChan <- rpc.NewClient(conn)
		}
	}()

	doClientWork3(clientChan)
}
