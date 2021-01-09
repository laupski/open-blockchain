package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/laupski/open-blockchain/api/proto"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartApi() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to listen to port 8080: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &server{
		Blockchain: blockchain.NewBlockChain(2,100.0),
	})
	fmt.Println("Running on port 8080, contact me at localhost:8080")
	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}

type server struct {
	Blockchain *blockchain.BlockChain
	proto.UnimplementedBlockchainServer
}

func (s *server) GetBlockchain(ctx context.Context, request *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)


	resp.Blockchain = &proto.BlockChain{
		Difficulty: s.Blockchain.Difficulty,
		MiningReward: s.Blockchain.MiningReward,
	}

	for _, b := range s.Blockchain.Chain {
		resp.Blockchain.Blocks = append(resp.Blockchain.Blocks, &proto.Block{
			Hash: base64.StdEncoding.EncodeToString(b.Hash),
			PrevHash: base64.StdEncoding.EncodeToString(b.PreviousHash),
		})
		/*
		for _, t := range b.Transactions {

		}*/
	}

	for _, t := range s.Blockchain.PendingTransactions {
		resp.Blockchain.PendingTransactions = append(resp.Blockchain.PendingTransactions, &proto.Transaction{
			FromAddress: base64.StdEncoding.EncodeToString(t.FromAddress),
			ToAddress: base64.StdEncoding.EncodeToString(t.ToAddress),
			Amount: t.Amount,
			Signature: base64.StdEncoding.EncodeToString(t.Signature),
		})
	}

	return resp, nil
}
