package main

import (
	"context"
	"encoding/hex"
	pb "github.com/alonmuroch/validatorremotewallet/wallet/v1alpha1"
	"google.golang.org/grpc"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.RemoteWalletServer
	conn         *grpc.Server
	port         string
	keys		[2]string
}

// SayHello implements helloworld.GreeterServer
func (s *Server) FetchValidatingKeys(ctx context.Context, in *pb.FetchValidatingKeysRequest) (*pb.FetchValidatingKeysResponse, error) {
	log.Printf("Received: get pub keys request",)

	// prepare payload
	payload := make([][]byte, len(s.keys))
	for index, key := range s.keys {
		payload[index], _ = hex.DecodeString(key)
	}

	return &pb.FetchValidatingKeysResponse{PublicKeys: payload}, nil
}

func (s *Server) start() (error) {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.conn = grpc.NewServer()
	pb.RegisterRemoteWalletServer(s.conn, s)

	log.Printf("Server up on port " + s.port)

	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
	return s.conn.Serve(lis)
}