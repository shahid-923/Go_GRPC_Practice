package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shahid-923/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send the names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending: %v", err)
		}
		log.Printf("Sent the request with name: %v", name)
		time.Sleep(2 * time.Second)
	}
	resp, err := stream.CloseAndRecv()
	log.Printf("Streaming finished")
	if err != nil {
		log.Fatalf("Error while recieving: %v", err)
	}
	log.Printf("%v", resp.Message)
}
