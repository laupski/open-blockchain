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

	createTransactionCmd.Flags().StringVarP(&toAddress, "toAddress", "t", "", "The address to send the amount (public key)")
	createTransactionCmd.Flags().Float32VarP(&amount, "amount", "a", 0, "Amount to send to the address")
	createTransactionCmd.Flags().StringVarP(&key, "key", "k", "", "The key to sign the transaction with")
	_ = createTransactionCmd.MarkFlagRequired("toAddress")
	_ = createTransactionCmd.MarkFlagRequired("amount")
	_ = createTransactionCmd.MarkFlagRequired("key")

	//sendTransactionCmd.Flags().StringVarP(&blockchainAddress, "blockchainAddress", "b", "", "The blockchain server address you want to send the transaction to")
	//_ = sendTransactionCmd.MarkFlagRequired("blockchainAddress")
}

var toAddress string
var amount float32
var key string
var createTransactionCmd = &cobra.Command{
	Use:   "create",
	Short: "Create transactions to send the blockchain server.",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := blockchain.LoadKey(key)
		if err != nil {
			fmt.Printf("Unable to load the key: %v", err)
			return
		}

		fromAddress := crypto.FromECDSAPub(key.Key.Public().(*ecdsa.PublicKey))
		toAddress, _ := base64.StdEncoding.DecodeString(toAddress)

		t := blockchain.NewTransaction(fromAddress, toAddress, amount)
		err = t.SignTransaction(key)
		if err != nil {
			fmt.Printf("Unable to sign the transaction: %v", err)
			return
		}

		err = t.SaveTransactionToJSON()
		if err != nil {
			fmt.Printf("Unable to save the transaction locally: %v", err)
			return
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
			fmt.Printf("Could not open transaction: %v", err)
			return
		}

		b := t.VerifyTransaction()
		if b != true {
			fmt.Printf("Transaction is NOT verified, did not send to blockchainn server")
			return
		}

		resp, err := client.SendTransaction(*t)
		if err != nil {
			fmt.Printf("Error with sending the transaction: %v", err)
			return
		}

		if resp.Confirmation == false {
			fmt.Println("Unknown error has occurred!")
		}

		fmt.Println("Transaction was successfully sent to the pending transaction list!")
	},
}

var printTransactionCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the transaction you created.",
	Run: func(cmd *cobra.Command, args []string) {
		t, err := blockchain.ReadTransactionFromJSON()
		if err != nil {
			fmt.Printf("could not open transaction: %v", err)
			return
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
			fmt.Printf("could not open transaction: %v\n", err)
			return
		}

		b := t.VerifyTransaction()
		if b == true {
			fmt.Printf("transaction is verified!\n")
		} else {
			fmt.Printf("transaction is NOT verified\n")
		}
	},
}
