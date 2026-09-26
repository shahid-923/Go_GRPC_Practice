package main

import (
	"context"
	"io"
	"log"

	pb "github.com/shahid-923/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started:")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send the names:%v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming%v", err)
		}
		log.Printf(message.Message)
	}
	log.Printf("Streaming finished")
}
