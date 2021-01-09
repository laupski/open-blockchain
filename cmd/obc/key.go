package obc

import (
	"github.com/laupski/open-blockchain/internal/blockchain"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Create and load keys.",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = blockchain.NewKey(args[0])
	},
}
