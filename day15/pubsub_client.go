package main

import (
	"context"
	"go-learn/day15/mypubsub"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := mypubsub.NewPubsubServiceClient(conn)

	_, err = client.Publish(
		context.Background(), &mypubsub.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &mypubsub.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
}
