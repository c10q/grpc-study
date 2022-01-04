package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	mailpb "grpc-study/email"
)

const portNum = "50001"

type mailServer struct {
	mailpb.MailServer
}

func (s *mailServer) Send(ctx context.Context, req *mailpb.SendRequest) (*mailpb.SendReply, error) {
	log.Println(req.To)
	log.Println(req.From)
	log.Println(req.Content)

	return &mailpb.SendReply{
		Success: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNum)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	mailpb.RegisterMailServer(grpcServer, &mailServer{})

	log.Printf("gRPC server on %s", portNum)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
