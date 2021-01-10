package obc

import (
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/laupski/open-blockchain/api/client"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"github.com/spf13/cobra"
)

var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Create and send transactions to the blockchain server.",
}

func init() {
	transactionCmd.AddCommand(createTransactionCmd)
	transactionCmd.AddCommand(sendTransactionCmd)
	transactionCmd.AddCommand(printTransactionCmd)
	transactionCmd.AddCommand(verifyTransactionCmd)

	createTransactionCmd.Flags().StringVarP(&ToAddress, "toAddress", "t", "", "The address to send the amount (public key)")
	createTransactionCmd.Flags().Float32VarP(&Amount, "amount", "a", 0, "Amount to send to the address")
	createTransactionCmd.Flags().StringVarP(&Key, "key", "k", "", "The key to sign the transaction with")
	_ = createTransactionCmd.MarkFlagRequired("toAddress")
	_ = createTransactionCmd.MarkFlagRequired("amount")
	_ = createTransactionCmd.MarkFlagRequired("key")

	//sendTransactionCmd.Flags().StringVarP(&blockchainAddress, "blockchainAddress", "b", "", "The blockchain server address you want to send the transaction to")
	//_ = sendTransactionCmd.MarkFlagRequired("blockchainAddress")
}

var ToAddress string
var Amount float32
var Key string
var createTransactionCmd = &cobra.Command{
	Use:   "create",
	Short: "Create transactions to send the blockchain server.",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := blockchain.LoadKey(Key)
		if err != nil {
			fmt.Printf("unable to load the key: %v", err)
		}

		fromAddress := crypto.FromECDSAPub(key.Key.Public().(*ecdsa.PublicKey))
		toAddress, _ := base64.StdEncoding.DecodeString(ToAddress)

		t := blockchain.NewTransaction(fromAddress, toAddress, Amount)
		err = t.SignTransaction(key)
		if err != nil {
			fmt.Printf("unable to sign the transaction: %v", err)
		}

		err = t.SaveTransactionToJSON()
		if err != nil {
			fmt.Printf("unable to save the transaction locally: %v", err)
		}
	},
}

//var blockchainAddress string
var sendTransactionCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a transaction to the blockchain server.",
	Run: func(cmd *cobra.Command, args []string) {
		t, err := blockchain.ReadTransactionFromJSON()
		if err != nil {
			fmt.Printf("could not open transaction: %v", err)
		}

		b := t.VerifyTransaction()
		if b != true {
			fmt.Printf("transaction is NOT verified, did not send to blockchainn server")
		}

		client.SendTransaction(*t)
	},
}

var printTransactionCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the transaction you created.",
	Run: func(cmd *cobra.Command, args []string) {
		t, err := blockchain.ReadTransactionFromJSON()
		if err != nil {
			fmt.Printf("could not open transaction: %v", err)
		}

		fmt.Println(t)
	},
}

var verifyTransactionCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the transaction you created.",
	Run: func(cmd *cobra.Command, args []string) {
		t, err := blockchain.ReadTransactionFromJSON()
		if err != nil {
			fmt.Printf("could not open transaction: %v", err)
		}

		b := t.VerifyTransaction()
		if b == true {
			fmt.Printf("transaction is verified!")
		} else {
			fmt.Printf("transaction is NOT verified")
		}
	},
}
