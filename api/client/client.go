package client

import (
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/laupski/open-blockchain/api/proto"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func GetBlockchain() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("cannot dial server: %v", err)
		return
	}
	defer conn.Close()

	client = proto.NewBlockchainClient(conn)

	response, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		fmt.Printf("unable to get blockchain: %v", err)
	}

	marshaller := jsonpb.Marshaler{
		EnumsAsInts: false,
		EmitDefaults: true,
		Indent: "  ",
		OrigName: true,
	}

	json,_ := marshaller.MarshalToString(response)
	fmt.Println(json)
}

func SendTransaction(t blockchain.Transaction) error {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
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
	resp, err := client.SendTransaction(context.Background(), &str)
	if err != nil {
		return err
	}
	if resp.Confirmation == true {
		fmt.Println("Successfully sent response to the server and added to pending transaction list!")
	} else {
		fmt.Printf("Unsuccessful attempt to send the transaction to the server: %s", resp.Message)
	}
	return nil
}

func MineBlock(address string) error {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client = proto.NewBlockchainClient(conn)
	mbr := proto.MineBlockRequest{Address: address}
	_, err = client.MineBlock(context.Background(), &mbr)
	return err
}

func VerifyBlockchain() error {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client = proto.NewBlockchainClient(conn)
	resp, err := client.VerifyBlockchain(context.Background(), &proto.VerifyBlockchainRequest{})
	if err != nil {
		fmt.Printf("could not verify ")
	}

	fmt.Println(resp.Verified)
	return err
}
