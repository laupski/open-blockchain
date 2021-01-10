package client

import (
	"fmt"
	"github.com/laupski/open-blockchain/api/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var client proto.BlockchainClient

func GetBlockchain() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
		return
	}

	client = proto.NewBlockchainClient(conn)

	blockchain, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("unable to get blockchain: %v", err)
	}

	fmt.Println(blockchain)
}

func addTransaction() {

	//confirmation, err := client.SendTransaction(context.Background(), )
}
