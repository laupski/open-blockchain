package obc

import (
	"github.com/laupski/open-blockchain/internal/blockchain"
	"github.com/spf13/cobra"
	"log"
)

var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Create and send transactions to the blockchain server.",
}

func init() {
	transactionCmd.AddCommand(createTransactionCmd)
	transactionCmd.AddCommand(sendTransactionCmd)
	transactionCmd.AddCommand(printTransactionCmd)

	createTransactionCmd.Flags().StringVarP(&ToAddress, "toAddress", "t", "", "The address to send the amount (public key)")
	createTransactionCmd.Flags().Float32VarP(&Amount,"amount", "a", 0,"Amount to send to the address")
	createTransactionCmd.Flags().StringVarP(&Key, "key", "k", "", "The key to sign the transaction with")
	_ = createTransactionCmd.MarkFlagRequired("toAddress")
	_ = createTransactionCmd.MarkFlagRequired("amount")
	_ = createTransactionCmd.MarkFlagRequired("key")

	sendTransactionCmd.Flags().StringVarP(&blockchainAddress, "blockchainAddress", "b", "", "The blockchain server address you want to send the transaction to")
	_ = sendTransactionCmd.MarkFlagRequired("blockchainAddress")
}

var ToAddress string
var Amount float32
var Key string
var createTransactionCmd = &cobra.Command{
	Use:   "create",
	Short: "Create transactions to send the blockchain server.",
	Run: func(cmd *cobra.Command, args []string) {
		_,err := blockchain.LoadKey(Key)
		if err != nil {
			log.Fatalf("unable to load the key: %v", err)
		}



	},
}

var blockchainAddress string
var sendTransactionCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a transaction to the blockchain server.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var printTransactionCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the transaction you created.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}