package obc

import (
	"fmt"
	"github.com/laupski/open-blockchain/api/server"
	"github.com/spf13/cobra"
	"strconv"
)

var port int
var difficulty int
var reward float32
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the blockchain server.",
	Run: func(cmd *cobra.Command, args []string) {
		if port < 1 {
			fmt.Println("Invalid port number!")
			return
		}

		if difficulty < 0 || difficulty > 10 {
			fmt.Println("Invalid difficulty selection!")
			return
		}

		if reward < 0 {
			fmt.Println("Invalid mining reward!")
			return
		}

		p := strconv.Itoa(port)
		server.StartAPI(p, difficulty, reward)
	},
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to start the blockchain server on.")
	serverCmd.Flags().IntVarP(&difficulty, "difficulty", "d", 2, "Difficulty to start the blockchain with.")
	serverCmd.Flags().Float32VarP(&reward, "reward", "r", 100.0, "Reward to set the blockchain with.")
	_ = serverCmd.MarkFlagRequired("port")
}
