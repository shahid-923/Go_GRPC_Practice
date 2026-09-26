package main

import (
	"log"
	"time"

	pb "github.com/shahid-923/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList,stream pb.GreetService_SayHelloServerStreamingServer) error {

	// Receive the list of names from the client
	log.Printf("Received names: %v", req.Names)

	// Send one response for each name
	for _, name := range req.Names {

		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		// Send the response to the client
		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}
	return nil
}