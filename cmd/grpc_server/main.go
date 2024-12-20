package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	desc "github.com/merynayr/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 8089

type server struct {
	desc.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to losten: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %s", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	log.Printf("Chat req: %v", req.Usernames)

	return &desc.CreateChatResponse{
		Id: gofakeit.Int64(),
	}, nil
}
