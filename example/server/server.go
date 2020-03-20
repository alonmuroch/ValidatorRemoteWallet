package main

import (
	"context"
	"fmt"
	pb "github.com/alonmuroch/validatorremotewallet/wallet/v1alpha1"
	"github.com/prysmaticlabs/go-ssz"
	bytesutil "github.com/wealdtech/go-bytesutil"
	types "github.com/wealdtech/go-eth2-wallet-types"
	"google.golang.org/grpc"
	"log"
	"net"
)

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
func (s *Server) SignAttestation(ctx context.Context, in *pb.SignAttestationRequest) (*pb.SignResponse, error) {
	log.Printf("Received: sign att. request")

	// create data to sign
	root, err := ssz.HashTreeRoot(in.Data)
	if err != nil {
		return nil, err
	}

	// find account
	account, exists := s.accountsMap[bytesutil.ToBytes48(in.PublicKey)]
	if !exists {
		return nil, fmt.Errorf("account does not exist")
	}

	// sign
	sig, err := sign(account, root, in.Domain)
	if err != nil {
		return nil, err
	}

	return &pb.SignResponse{
		Sig: sig,
		Status: pb.SignResponseStatus_SIGNED,
	}, nil
}

func (s *Server) SignProposal(ctx context.Context, in *pb.SignProposalRequest) (*pb.SignResponse, error) {
	log.Printf("Received: sign proposal request")

	// create data to sign
	blockRoot, err := ssz.HashTreeRoot(in.Data)
	if err != nil {
		return nil, err
	}

	// find account
	account, exists := s.accountsMap[bytesutil.ToBytes48(in.PublicKey)]
	if !exists {
		return nil, fmt.Errorf("account does not exist")
	}

	// sign
	sig, err := sign(account, blockRoot, in.Domain)
	if err != nil {
		return nil, err
	}

	return &pb.SignResponse{
		Sig: sig,
		Status: pb.SignResponseStatus_SIGNED,
	}, nil
}

func (s *Server) SignSlot(ctx context.Context, in *pb.SignSlotRequest) (*pb.SignResponse, error) {
	log.Printf("Received: sign slot request")

	// create data to sign
	slotRoot, err := ssz.HashTreeRoot(in.Slot)
	if err != nil {
		return nil, err
	}

	// find account
	account, exists := s.accountsMap[bytesutil.ToBytes48(in.PublicKey)]
	if !exists {
		return nil, fmt.Errorf("account does not exist")
	}

	// sign
	sig, err := sign(account, slotRoot, in.Domain)
	if err != nil {
		return nil, err
	}

	return &pb.SignResponse{
		Sig: sig,
		Status: pb.SignResponseStatus_SIGNED,
	}, nil
}

func sign(account types.Account, toSign [32]byte, domain uint64) ([]byte, error) {
	sig, err := account.Sign(toSign[:], domain)
	if err != nil {
		return nil, err
	}
	return sig.Marshal(), nil
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