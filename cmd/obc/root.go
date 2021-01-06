package obc

import (
	"fmt"
	"github.com/laupski/open-blockchain/internal/blockchain"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "obc",
	Short: "open-blockchain is a PoC blockchain generator",
	Long: `A simple proof of concept of blockchain technology 
                written by laupski in Go.
                Complete documentation is available at https://github.com/laupski/open-blockchain`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO - placeholders for now until persistence
		fmt.Println("Creating Blockchain")
		bc := blockchain.NewBlockChain("MyBlockchain")
		fmt.Println(bc)

		bc.AppendBlock(blockchain.NewBlock("2nd Block", bc.GetLatestBlock().Hash))
		bc.AppendBlock(blockchain.NewBlock("3rd Block", bc.GetLatestBlock().Hash))
		bc.AppendBlock(blockchain.NewBlock("4th Block", bc.GetLatestBlock().Hash))
		fmt.Println(bc)

		fmt.Println("Verifying Blockchain")
		fmt.Println(bc.Verify())
	},
}

// Execute runs the command line input.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
