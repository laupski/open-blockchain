package obc

import (
	"fmt"
	"github.com/laupski/open-blockchain/api/client"
	"github.com/spf13/cobra"
)

var valid = map[string]func() {
	"blockchain" : client.GetBlockchain,
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information from a blockchain server.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("no arguments supplied")
			return
		} else if len(args) > 1 {
			fmt.Println("too many arguments supplied")
			return
		}

		function, ok := valid[args[0]]
		if ok != true {
			fmt.Println("not a valid option under get")
		} else {
			function()
		}
	},
}
