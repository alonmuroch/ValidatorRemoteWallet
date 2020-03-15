package main

import (
	"context"
	"encoding/hex"
	pb "github.com/alonmuroch/validatorremotewallet/wallet/v1alpha1"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
)


func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRemoteWalletClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.FetchValidatingKeys(ctx, &pb.FetchValidatingKeysRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for _, key := range r.PublicKeys {
		log.Printf("key: %s", hex.EncodeToString(key))
	}

}

