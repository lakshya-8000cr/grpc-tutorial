# gRPC Practice (Go)

Learning gRPC from scratch using Go.

## Goal

Understand how gRPC works internally instead of only writing code.

Topics covered:

- Protocol Buffers (.proto)
- gRPC services
- Request/Response messages
- Code generation with protoc
- gRPC server
- gRPC client
- HTTP/2 basics
- Binary serialization

## Project Structure

```text
grpc-practice/

├── client/
│   └── main.go

├── server/
│   └── main.go

├── proto/
│   └── hello.proto

├── grpc-practice/
│   └── proto/
│       └── hellopb/
│           ├── hello.pb.go
│           └── hello_grpc.pb.go

├── go.mod

└── README.md
```

## Architecture

```text
Client

↓

GreeterClient

↓

HTTP/2

↓

gRPC Server

↓

SayHello()

↓

HelloResponse
```

## Commands

Generate protobuf code:

```bash
protoc --go_out=. --go-grpc_out=. proto/hello.proto
```

Run server:

```bash
go run server/main.go
```

Run client:

```bash
go run client/main.go
```
