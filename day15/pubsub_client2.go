package main

import (
	"context"
	"fmt"
	"io"
	"log"

	mypubsub "./mypubsub"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := mypubsub.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(
		context.Background(), &mypubsub.String{Value: "golang:"},
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}
