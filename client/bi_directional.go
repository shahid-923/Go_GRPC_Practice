package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/shahid-923/proto"
)

func callHelloBidrectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Bidirectional streaming started")

	stream, err := client.SayHelloBidrectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not start stream: %v", err)
	}

	wait := make(chan struct{})

	// Receiving responses from server
	go func() {
		for {
			message, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while streaming: %v", err)
			}

			log.Printf("Received: %s", message.Message)
		}

		close(wait)
	}()

	// Sending requests to server
	for _, name := range names.Names {

		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}

		log.Printf("Sent: %s", name)

		time.Sleep(2 * time.Second)
	}

	// Tell server that client has finished sending
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error closing stream: %v", err)
	}

	// Wait until all responses are received
	<-wait

	log.Println("Bidirectional streaming finished")
}
