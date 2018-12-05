package main

import (
	"fmt"
	"github.com/matthieuberger/go-blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Data 1 stored in the blockchain")
	bc.AddBlock("Data 2 stored in the blockchain")

	for _, block := range bc.Blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
