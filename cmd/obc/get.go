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
		response, err := client.GetBlockchain()
		if err != nil {
			fmt.Printf("Could not contact server: %v\n", err)
			return
		}

		json, err := client.Marshaller.MarshalToString(response)
		if err != nil {
			fmt.Printf("Error on unmarshalling: %v\n", err)
			return
		}

		fmt.Println(json)
	},
}

var verifyGetCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the current blockchain.",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := client.VerifyBlockchain()
		if err != nil {
			fmt.Printf("could not verify the blockchain: %v\n", err)
			return
		}

		fmt.Printf("Blockchain verified: %v\n", resp.Verified)
	},
}
