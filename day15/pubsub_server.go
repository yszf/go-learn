package main

import (
	"go-learn/day15/mypubsub"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	mypubsub.RegisterPubsubServiceServer(grpcServer, mypubsub.NewPubsubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)
}
