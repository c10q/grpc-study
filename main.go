package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

const portNum = "50001"

func main() {
	lis, err := net.Listen("tcp", ":"+portNum)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("gRPC server on %s", portNum)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
