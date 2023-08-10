package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc/grpc/proto"
	"log"
	"net"
)

var port = "localhost:3000"

type server struct {
	proto.MessangerServer
}

func (s *server) Messager(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	mes := "hello, " + req.Message + " good man"
	return &proto.Response{
		Message: mes,
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterMessangerServer(s, &server{})

	log.Printf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
