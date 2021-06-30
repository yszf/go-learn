package main

import (
	"go-learn/day15/kv"
	"log"
	"net"
	"net/rpc"
)

func main() {
	err := rpc.RegisterName("KVStoreService", kv.NewKVStoreService())
	if err != nil {
		return
	}

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
