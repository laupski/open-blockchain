package obc

import (
	"fmt"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"github.com/spf13/cobra"
	"log"
)

var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Prints your address from a supplied key.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("must supply a key name")
		}

		address, err := blockchain.PrintPublicAddress(args[0])
		if err != nil {
			log.Fatalf("unable to read the address of the key: %v", err)
		}
		fmt.Printf("Address of %s: %s", args[0], address)
	},
}

func init() {

}
