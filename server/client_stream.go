package main

import (
	"io"
	"log"

	pb "github.com/shahid-923/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv() // Keep receiving until client finishes
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{
				Message: messages,
			})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name:%v", req.Name)
		messages = append(messages, "Hello:", req.Name)
	}
}
