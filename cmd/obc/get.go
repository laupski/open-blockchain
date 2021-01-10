package obc

import (
	"fmt"
	"github.com/laupski/open-blockchain/api/client"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information from a blockchain server.",
}

func init() {
	getCmd.AddCommand(blockchainGetCmd)
	getCmd.AddCommand(verifyGetCmd)
}

var blockchainGetCmd = &cobra.Command{
	Use:   "blockchain",
	Short: "Get the current blockchain.",
	Run: func(cmd *cobra.Command, args []string) {
		client.GetBlockchain()
	},
}

var verifyGetCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the current blockchain.",
	Run: func(cmd *cobra.Command, args []string) {
		err := client.VerifyBlockchain()
		if err != nil {
			fmt.Printf("could not verify the blockchain: %v", err)
		}
	},
}
