package client

import (
	"encoding/base64"
	"github.com/golang/protobuf/jsonpb"
	"github.com/laupski/open-blockchain/api/proto"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Marshaller default for the client.
var Marshaller = jsonpb.Marshaler{
	EnumsAsInts:  false,
	EmitDefaults: true,
	Indent:       "  ",
	OrigName:     true,
}
var client proto.BlockchainClient

// GetBlockchain sends a request to the node server for its Blockchain
func GetBlockchain() (*proto.GetBlockchainResponse, error) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client = proto.NewBlockchainClient(conn)

	return client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
}

// SendTransaction sends a transaction to a node server to add it to its pending transactions
func SendTransaction(t blockchain.Transaction) (*proto.SendTransactionResponse, error) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client = proto.NewBlockchainClient(conn)
	pt := proto.Transaction{
		FromAddress: base64.StdEncoding.EncodeToString(t.FromAddress),
		ToAddress:   base64.StdEncoding.EncodeToString(t.ToAddress),
		Amount:      t.Amount,
		Signature:   base64.StdEncoding.EncodeToString(t.Signature),
	}
	str := proto.SendTransactionRequest{Transaction: &pt}
	return client.SendTransaction(context.Background(), &str)
}

// MineBlock sends a request to the node to mine the next block
func MineBlock(address string) (*proto.MineBlockResponse, error) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client = proto.NewBlockchainClient(conn)
	mbr := proto.MineBlockRequest{Address: address}
	return client.MineBlock(context.Background(), &mbr)
}

// VerifyBlockchain sends a request to the node to verify its blockchain
func VerifyBlockchain() (*proto.VerifyBlockchainResponse, error) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client = proto.NewBlockchainClient(conn)
	return client.VerifyBlockchain(context.Background(), &proto.VerifyBlockchainRequest{})
}
