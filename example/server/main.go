package main

import (
	pb "../../wallet/v1alpha1"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedRemoteWalletServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) FetchValidatingKeys(ctx context.Context, in *pb.FetchValidatingKeysRequest) (*pb.FetchValidatingKeysResponse, error) {
	log.Printf("Received: get pub keys request",)
	return &pb.FetchValidatingKeysResponse{PublicKeys: [][]byte{[]byte("pub1")}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRemoteWalletServer(s, &server{})

	log.Printf("Server up on port " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
