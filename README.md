# Go gRPC Streaming Practice

A beginner-friendly Go project for learning **gRPC communication and different types of RPC streaming** using Protocol Buffers.

## Project Structure

GO_GRPC/
├── client/
│   └── main.go
├── server/
│   ├── main.go
│   └── ...
├── proto/
│   ├── greet.proto
│   ├── greet.pb.go
│   └── greet_grpc.pb.go
├── go.mod
└── README.md
````

## Technologies

* Go
* gRPC
* Protocol Buffers (Protobuf)
* grpc-go

## gRPC Communication Types

This project covers all four major gRPC communication patterns.

### 1. Unary RPC

One request → One response

```text
Client ── Request ──> Server
Client <── Response ── Server
```

Example:

```proto
rpc SayHello(NoParam) returns (HelloResponse);
```

---

### 2. Server Streaming

One request → Multiple responses

Client ── Request ──> Server

Client <── Response ── Server
Client <── Response ── Server
Client <── Response ── Server
```

Example:

```proto
rpc SayHelloServerStreaming(NamesList)
    returns (stream HelloResponse);
```

The server uses:

```go
stream.Send(res)
```

The client receives multiple responses using:

```go
stream.Recv()
```

---

### 3. Client Streaming

Multiple requests → One response

```text
Client ── Request ──> Server
Client ── Request ──> Server
Client ── Request ──> Server

Client <── Response ── Server
```

Example:

```proto
rpc SayHelloClientStreaming(stream HelloRequest)
    returns (MessageList);
```

The client sends multiple requests:

```go
stream.Send(req)
```

After sending all requests:

```go
stream.CloseAndRecv()
```

The server receives requests using:

```go
stream.Recv()
```

and sends the final response using:

```go
stream.SendAndClose(...)
```

---

### 4. Bidirectional Streaming

Multiple requests ↔ Multiple responses

```text
Client ── Request ─────> Server
Client <── Response ──── Server
Client ── Request ─────> Server
Client <── Response ──── Server
```

Example:

```proto
rpc SayHelloBidrectionalStreaming(stream HelloRequest)
    returns (stream HelloResponse);
```

The client can send and receive messages independently.

The client uses:

```go
stream.Send(req)
stream.Recv()
```

Usually, receiving is handled in a separate goroutine while requests are being sent.

## Proto Definition

The main `.proto` file defines the messages and gRPC service.

Example:

```proto
syntax = "proto3";

option go_package = "./proto";

package greet_service;

service GreetService {
  rpc SayHello(NoParam) returns (HelloResponse);

  rpc SayHelloServerStreaming(NamesList)
      returns (stream HelloResponse);

  rpc SayHelloClientStreaming(stream HelloRequest)
      returns (MessageList);

  rpc SayHelloBidrectionalStreaming(stream HelloRequest)
      returns (stream HelloResponse);
}

message NoParam {}

message HelloRequest {
  string name = 1;
}

message NamesList {
  repeated string names = 1;
}

message HelloResponse {
  string message = 1;
}

message MessageList {
  repeated string message = 1;
}
```

## Generate Go Code

From the project root:

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

This generates:

```text
greet.pb.go
greet_grpc.pb.go
```

### What are these files?

`greet.pb.go`

Contains generated Go structures for Protobuf messages such as:

```go
pb.HelloRequest
pb.HelloResponse
pb.NamesList
pb.MessageList
```

`greet_grpc.pb.go`

Contains generated gRPC client/server code such as:

```go
pb.NewGreetServiceClient()
pb.RegisterGreetServiceServer()
```

## What is `pb`?

In the Go code:

```go
import pb "github.com/shahid-923/proto"
```

`pb` is simply an **import alias** for the generated protobuf package.

For example:

```go
req := &pb.HelloRequest{
    Name: "Shahid",
}
```

The `HelloRequest` type was generated from:

```proto
message HelloRequest {
    string name = 1;
}
```

## Running the Project

### Start the server

Open a terminal:

```bash
go run ./server
```

The server listens on:

```text
localhost:8080
```

### Start the client

Open another terminal:

```bash
go run ./client
```

Make sure the server is running before starting the client.

## Learning Flow

The recommended order for understanding this project is:

```text
1. Protobuf
   ↓
2. Unary RPC
   ↓
3. Server Streaming
   ↓
4. Client Streaming
   ↓
5. Bidirectional Streaming
```

The main concepts practiced are:

* `.proto` files
* Protocol Buffers
* Generated Go code
* gRPC client
* gRPC server
* RPC methods
* `stream.Send()`
* `stream.Recv()`
* `CloseAndRecv()`
* `SendAndClose()`
* `io.EOF`
* Goroutines
* Bidirectional communication

## Important Concept

The `.proto` file is the **contract** between the client and server.

```text
             greet.proto
                  │
             protoc compiler
                  │
        ┌─────────┴─────────┐
        ▼                   ▼
   greet.pb.go        greet_grpc.pb.go
        │                   │
        └─────────┬─────────┘
                  ▼
          Go Client & Server
```

Both the client and server use the generated code, so they agree on the message structures and RPC methods.

```
```
