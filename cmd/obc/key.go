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
		_, err := blockchain.NewKey(args[0])
		if err != nil {
			fmt.Printf("unable to create key: %v", err)
		}
	},
}
