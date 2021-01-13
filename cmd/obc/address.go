package obc

import (
	"fmt"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"github.com/spf13/cobra"
)

var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Prints your address from a supplied key.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Must supply a key name")
			return
		}

		address, err := blockchain.PrintPublicAddress(args[0])
		if err != nil {
			fmt.Printf("Unable to read the address of the key: %v\n", err)
			return
		}
		fmt.Println(address)
	},
}
