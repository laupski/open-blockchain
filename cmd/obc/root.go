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
		bc := blockchain.NewBlockChain(4, 100)

		txl := make(blockchain.TransactionList, 0)
		txl = append(txl, blockchain.NewTransaction("address 1", "address 2", 100))
		txl = append(txl, blockchain.NewTransaction("address 1", "address 3", 50))
		txl = append(txl, blockchain.NewTransaction("address 1", "address 4", 30))
		txl = append(txl, blockchain.NewTransaction("address 1", "address 5", 20))

		bc.PushTransactions(txl...)
		bc.MineTransactions("home")

		fmt.Printf("Balance of %v: %v\n", "address 1", bc.GetBalance("address 1"))
		fmt.Printf("Balance of %v: %v\n", "address 2", bc.GetBalance("address 2"))
		fmt.Printf("Balance of %v: %v\n", "address 3", bc.GetBalance("address 3"))
		fmt.Printf("Balance of %v: %v\n", "address 4", bc.GetBalance("address 4"))
		fmt.Printf("Balance of %v: %v\n", "address 5", bc.GetBalance("address 5"))
		fmt.Printf("Balance of %v: %v\n", "home", bc.GetBalance("home"))

		fmt.Println(bc)
		fmt.Println("Verifying Blockchain")
		fmt.Println(bc.Verify())

		txl = blockchain.TransactionList{blockchain.NewTransaction("address 1", "home", 100)}
		bc.PushTransactions(txl...)
		bc.MineTransactions("home")

		fmt.Println()
		fmt.Printf("Balance of %v: %v\n", "home", bc.GetBalance("home"))
		fmt.Printf("Balance of %v: %v\n", "address 1", bc.GetBalance("address 1"))
		fmt.Println(bc)
		fmt.Println("Verifying Blockchain")
		fmt.Println(bc.Verify())

		// TODO write GetAllBalances()
	},
}

// Execute runs the command line input.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
