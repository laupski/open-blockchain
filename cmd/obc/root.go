package obc

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "obc",
	Short: "open-blockchain is a PoC blockchain generator",
	Long: `A simple proof of concept of blockchain technology 
                written by laupski in Go.
                Complete documentation is available at https://github.com/laupski/open-blockchain`,
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(keyCmd)
	rootCmd.AddCommand(transactionCmd)
	rootCmd.AddCommand(demoCmd)
	rootCmd.AddCommand(addressCmd)
}

// Execute runs the command line input.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
