package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Calculate a SHA256 hash for a Block.
func (b *Block) CalculateHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, b.Data, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// New Block generates a new Block
func NewBlock(prevBlockHash []byte, data string) *Block {
	block := &Block{Timestamp: time.Now().Unix(), PrevBlockHash: prevBlockHash, Data: []byte(data)}
	block.CalculateHash()
	return block
}

// Return the latest block recorded in the blockchain
func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

// Add a block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.GetLastBlock()
	newBlock := NewBlock(prevBlock.Hash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// Replace the Blockchain
func (bc *Blockchain) Replace(newBc *Blockchain) {
	// Replace the blockchain only if the new blockchain
	// length is greater than the current one
	if len(newBc.Blocks) > len(bc.Blocks) {
		bc.Blocks = newBc.Blocks
	}
}

// First Block in the chain is called Genesis Block
func NewGenesisBlock() *Block {
	return NewBlock([]byte{}, "Genesis Block")
}

// Create a Blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{NewGenesisBlock()}}
}
