package obc

import (
	"github.com/laupski/open-blockchain/api/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Describes server.",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartApi()
	},
}
