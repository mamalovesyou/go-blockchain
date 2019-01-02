package cmd

import (
	"github.com/matthieuberger/go-blockchain/blockchain"

	"github.com/matthieuberger/go-blockchain/server"
	"github.com/spf13/cobra"
)

var peers []string
var port int
var addr string
var protocol string

var host *server.P2PHost = nil

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new blockchain",
	Run: func(cmd *cobra.Command, args []string) {

		// Create a new blockchain
		bc := blockchain.NewBlockchain()
		plugin := &server.BlockChainServerPlugin{RemoteAddress: addr, Blockchain: bc}
		host := server.P2PHost{port, protocol, plugin, peers}
		host.Start()

	},
}

func init() {

	startCmd.Flags().StringSliceVar(&peers, "peers", []string{}, "Peers addresses")
	startCmd.Flags().IntVarP(&port, "port", "p", 8555, "Listening port")
	startCmd.Flags().StringVarP(&addr, "address", "a", "0.0.0.0", "Address")
	startCmd.Flags().StringVar(&protocol, "protocol", "tcp", "Protocol (Ex: tcp)")

	rootCmd.AddCommand(startCmd)
}
