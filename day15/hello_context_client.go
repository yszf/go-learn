package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func Login(client *rpc.Client) error {
	var reply string
	return client.Call("HelloService.Login", "user:password", &reply)
}

func Hello(client *rpc.Client) {
	defer client.Close()

	helloCall := client.Go("HelloService.Hello", "hello", new(string), nil)

	// do some thing
	helloCall = <-helloCall.Done
	if err := helloCall.Error; err != nil {
		log.Fatal(err)
	}

	//	args := helloCall.Args.(string)
	reply := helloCall.Reply.(*string)
	fmt.Println(*reply)
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	err = Login(client)
	if nil != err {
		log.Fatal("login:", err)
	}

	Hello(client)
}
