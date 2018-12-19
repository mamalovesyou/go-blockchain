package server

import (
	"context"
	"log"

	"github.com/matthieuberger/go-blockchain/blockchain"

	"github.com/perlin-network/noise/network"
)

type BlockChainServerPlugin struct {
	network.Plugin
	RemoteAddress string
	Blockchain    *blockchain.Blockchain
}

func (state *BlockChainServerPlugin) PeerConnect(client *network.PeerClient) {
	log.Printf("*** New connection from %s.\n", client.Address)

	// Send the blockchain to new peers
	client.Reply(context.Background(), 0, state.Blockchain)
}

func (state *BlockChainServerPlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *blockchain.Blockchain:
		// If the new blockchain is longer than the one we have then replace it
		if len(msg.Blocks) > len(state.Blockchain.Blocks) {
			log.Printf("[UPDATE BLOCKCHAIN] From: %s\n%v\n", ctx.Client().ID.Address, msg)
			state.Blockchain = msg
		}
	}
	return nil
}

func (state *BlockChainServerPlugin) PeerDisconnect(client *network.PeerClient) {
	log.Println("*** Lost connection with %s.", client.Address)
}
