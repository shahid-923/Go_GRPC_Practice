package main

import (
	"log"
	"net"

	pb "github.com/shahid-923/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer                                       // embeds the GreetServer interface we implement the RPC methods here lke SayHello
}

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server:%v", err)
	}

	grpcServer := grpc.NewServer()                              // creates a server instance

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})   // connects go code to generated protobuf service def

	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {               // begins accepting client connections @port
		log.Fatalf("Failed to start: %v", err)
	}
}
