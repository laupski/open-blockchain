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
		bc := blockchain.NewBlockChain("MyBlockchain", 4)

		bc.AppendBlock(blockchain.NewBlock("2nd"))
		bc.AppendBlock(blockchain.NewBlock("3rd"))
		bc.AppendBlock(blockchain.NewBlock("4th"))

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
