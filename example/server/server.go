package main

import (
	"context"
	"fmt"
	pb "github.com/alonmuroch/validatorremotewallet/wallet/v1alpha1"
	"github.com/wealdtech/go-bytesutil"
	types "github.com/wealdtech/go-eth2-wallet-types"
	"google.golang.org/grpc"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.RemoteWalletServer
	conn         	*grpc.Server
	port         	string
	accountsMap 	map[[48]byte]types.Account
}

func (s *Server) FetchValidatingKeys(ctx context.Context, in *pb.FetchValidatingKeysRequest) (*pb.FetchValidatingKeysResponse, error) {
	log.Printf("Received: get pub keys request")

	// prepare payload
	payload := make([][]byte, 0, len(s.accountsMap))
	for pubKey := range s.accountsMap {
		tmp := make([]byte, len(pubKey))
		copy(tmp, pubKey[:])
		payload = append(payload,  tmp)
	}

	return &pb.FetchValidatingKeysResponse{PublicKeys: payload}, nil
}

// Sign signs a message for the validator to broadcast.
func (s *Server) Sign(ctx context.Context, in *pb.SignRequest) (*pb.SignResponse, error) {
	//account, exists := km.accounts[pubKey]
	//if !exists {
	//	return nil, ErrNoSuchKey
	//}
	//sig, err := account.Sign(root[:], domain)
	//if err != nil {
	//	return nil, err
	//}
	//return bls.SignatureFromBytes(sig.Marshal())

	log.Printf("Received: sign reques")

	account, exists := s.accountsMap[bytesutil.ToBytes48(in.PublicKey)]
	if !exists {
		return nil, fmt.Errorf("account does not exist")
	}
	sig, err := account.Sign(in.Root[:], in.Domain)
	if err != nil {
		return nil, err
	}

	return &pb.SignResponse{
		Sig: sig.Marshal(),
		Status: pb.SignResponseStatus_SIGNED,
	}, nil
}

func (s *Server) start() (error) {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.conn = grpc.NewServer()
	pb.RegisterRemoteWalletServer(s.conn, s)

	log.Printf("Server up on port " + s.port)
	return s.conn.Serve(lis)
}