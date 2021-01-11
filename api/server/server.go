package server

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/laupski/open-blockchain/api/proto"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"google.golang.org/grpc"
	"log"
	"net"
)

var nodeKey *blockchain.Key

// StartAPI starts a listening gRPC server.
func StartAPI(port string, difficulty int, reward float32) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("unable to listen to port %s: %v\n", port, err)
	}

	nodeKey, _ = blockchain.NewKey("node.key")

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &server{
		Blockchain: blockchain.NewBlockChain(int32(difficulty), reward),
	})
	fmt.Printf("Creating blockchain with difficulty of %v and mining reward of %f\n", difficulty, reward)
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

// GetBlockchain sends back a response with the node's current entire Blockchain
func (s *server) GetBlockchain(ctx context.Context, request *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	log.Println("Received request for GetBlockchain")
	resp := new(proto.GetBlockchainResponse)

	resp.Blockchain = &proto.BlockChain{
		Difficulty:   s.Blockchain.Difficulty,
		MiningReward: s.Blockchain.MiningReward,
	}

	for _, b := range s.Blockchain.Chain {
		transaction := new(proto.Transaction)
		var tl []*proto.Transaction
		for _, t := range b.Transactions {
			transaction.Signature = base64.StdEncoding.EncodeToString(t.Signature)
			transaction.Amount = t.Amount
			transaction.ToAddress = base64.StdEncoding.EncodeToString(t.ToAddress)
			transaction.FromAddress = base64.StdEncoding.EncodeToString(t.FromAddress)
			tl = append(tl, transaction)
		}
		resp.Blockchain.Blocks = append(resp.Blockchain.Blocks, &proto.Block{
			Hash:         base64.StdEncoding.EncodeToString(b.Hash),
			PrevHash:     base64.StdEncoding.EncodeToString(b.PreviousHash),
			Transactions: tl,
		})
	}

	for _, t := range s.Blockchain.PendingTransactions {
		resp.Blockchain.PendingTransactions = append(resp.Blockchain.PendingTransactions, &proto.Transaction{
			FromAddress: base64.StdEncoding.EncodeToString(t.FromAddress),
			ToAddress:   base64.StdEncoding.EncodeToString(t.ToAddress),
			Amount:      t.Amount,
			Signature:   base64.StdEncoding.EncodeToString(t.Signature),
		})
	}

	return resp, nil
}

// SendTransaction accepts an incoming Transaction, verifies the signature and if validated, adds it to its pending Transaction list
func (s *server) SendTransaction(ctx context.Context, request *proto.SendTransactionRequest) (*proto.SendTransactionResponse, error) {
	log.Println("Received request for SendTransaction")
	resp := new(proto.SendTransactionResponse)
	t1 := blockchain.Transaction{}
	t1.FromAddress, _ = base64.StdEncoding.DecodeString(request.Transaction.FromAddress)
	t1.ToAddress, _ = base64.StdEncoding.DecodeString(request.Transaction.ToAddress)
	t1.Amount = request.Transaction.Amount
	t1.Signature, _ = base64.StdEncoding.DecodeString(request.Transaction.Signature)

	v := t1.VerifyTransaction()
	if v != true {
		resp.Confirmation = false
		resp.Message = "transaction could not be verified"
		return resp, nil
	}
	txl := make(blockchain.TransactionList, 0)
	txl = append(txl, t1)

	err := s.Blockchain.PushTransactions(txl...)
	if err != nil {
		resp.Confirmation = false
		resp.Message = err.Error()
		return resp, nil
	}

	resp.Confirmation = true
	resp.Message = ""
	return resp, nil
}

// MineBlock starts mining the next block. Credit is given to the owner of the node.
func (s *server) MineBlock(ctx context.Context, request *proto.MineBlockRequest) (*proto.MineBlockResponse, error) {
	log.Printf("Received request for MineBlock from Address: %s\n", request.Address)
	resp := new(proto.MineBlockResponse)

	// Even though we get a request from an address, the node will get credit for mining of the block
	s.Blockchain.MineTransactions(crypto.FromECDSAPub(nodeKey.Key.Public().(*ecdsa.PublicKey)))

	return resp, nil
}

// VerifyBlockchain verifies the integrity of the node's blockchain.
func (s *server) VerifyBlockchain(ctx context.Context, request *proto.VerifyBlockchainRequest) (*proto.VerifyBlockchainResponse, error) {
	log.Println("Received request for VerifyBlockchain")
	resp := new(proto.VerifyBlockchainResponse)

	resp.Verified = s.Blockchain.Verify()

	return resp, nil
}
