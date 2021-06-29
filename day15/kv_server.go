package main

import (
	"log"
	"net"
	"net/rpc"

	"./kv"
)

func main() {
	rpc.RegisterName("KVStoreService", kv.NewKVStoreService())

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
