package cmd

import (
	"github.com/matthieuberger/go-blockchain/server"
	"github.com/spf13/cobra"
	"os"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a nwe blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("GO_BLOCKCHAIN_TCP_ADDR")
		server := server.NewBlockChainServer(port)
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
