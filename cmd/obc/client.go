package obc

import (
	"github.com/laupski/open-blockchain/api/client"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Makes requests to a blockchain server.",
	Run: func(cmd *cobra.Command, args []string) {
		client.RunClient()
	},
}