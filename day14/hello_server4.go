package main

import (
	"log"
	"net"
	"net/rpc"

	"./hello"
)

type HelloService struct{}

func (p *HelloService) Hello(request *hello.String, reply *hello.String) error {
	(*reply).Value = "hello:" + request.GetValue()
	return nil
}

func main() {
	hello.RegisterHelloService(rpc.DefaultServer, new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
