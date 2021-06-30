package main

import (
	"net"
	"net/rpc"
	"time"
)

type HelloService2 struct{}

func (p *HelloService2) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.Register(new(HelloService2))

	for {
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}
}
