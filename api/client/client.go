package client

import (
	"fmt"
	"github.com/laupski/open-blockchain/api/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var client proto.BlockchainClient

func RunClient() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
		return
	}

	client = proto.NewBlockchainClient(conn)
	getBlockchain()
	fmt.Println("Successfully connected to the blockchain server.")
}

/*
func addBlock() {
	block, addErr := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if addErr != nil {
		log.Fatalf("unable to add block: %v", addErr)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}
*/
func getBlockchain() {
	blockchain, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("unable to get blockchain: %v", err)
	}

	fmt.Println(blockchain)
}
