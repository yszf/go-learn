package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	mypubsub "./mypubsub"
)

func main() {
	grpcServer := grpc.NewServer()
	mypubsub.RegisterPubsubServiceServer(grpcServer, (mypubsub.NewPubsubService()))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)
}
