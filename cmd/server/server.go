package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}