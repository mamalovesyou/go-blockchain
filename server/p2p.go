package server

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/matthieuberger/go-blockchain/blockchain"
	"github.com/perlin-network/noise/crypto/ed25519"
	"github.com/perlin-network/noise/network"
	"github.com/perlin-network/noise/network/discovery"
	"github.com/perlin-network/noise/types/opcode"
)

type P2PHost struct {
	Port     int
	Protocol string
	Plugin   *BlockChainServerPlugin
	Peers    []string
}

func (state *P2PHost) Start() {

	keys := ed25519.RandomKeyPair()

	opcode.RegisterMessageType(opcode.Opcode(1000), &blockchain.Blockchain{Blocks: []*blockchain.Block{}})
	opcode.RegisterMessageType(opcode.Opcode(2000), &blockchain.Block{})

	builder := network.NewBuilder()
	builder.SetKeys(keys)
	builder.SetAddress(network.FormatAddress(state.Protocol, state.Plugin.RemoteAddress, uint16(state.Port)))

	// Automatic peer discovery based off of a secure Kademlia-inspired routing table structure
	builder.AddPlugin(new(discovery.Plugin))
	builder.AddPlugin(state.Plugin)

	net, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	go net.Listen()

	if len(state.Peers) > 0 {
		net.Bootstrap(state.Peers...)
	}

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("*** > Enter a message: ")
		input, _ := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Broadcast the blockchain out to peers.
		state.Plugin.Blockchain.AddBlock(input)
		net.Broadcast(context.Background(), state.Plugin.Blockchain)
	}

}
