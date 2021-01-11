package obc

import (
	"fmt"
	"github.com/laupski/open-blockchain/api/client"
	"github.com/spf13/cobra"
)

var requestAddress string
var mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Send a request to mine a block on the server.",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := client.MineBlock(requestAddress)
		if err != nil {
			fmt.Printf("Could not send request to the server to mine a block: %v\n", err)
			return
		}

		fmt.Println("Successfully sent request to the server to mine a block!")
	},
}

func init() {
	mineCmd.Flags().StringVarP(&requestAddress, "requestAddress", "r", "", "Your address to send the request. Note: You will note receive credit for mining, the node will.")
	_ = mineCmd.MarkFlagRequired("requestAddress")
}