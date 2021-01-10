package obc

import (
	"github.com/laupski/open-blockchain/internal/blockchain"
	"github.com/spf13/cobra"
	"log"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Create and load keys.",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := blockchain.NewKey(args[0])
		if err != nil {
			log.Fatalf("unable to create key: %v", err)
		}
	},
}
