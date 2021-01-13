package obc

import (
	"fmt"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Create and load keys.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Need to supply key name in .key extension format")
			return
		}

		_, err := blockchain.NewKey(args[0])
		if err != nil {
			fmt.Printf("Unable to create key: %v\n", err)
			return
		}
	},
}
