package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Define Block struct
type Block struct {
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
}

// Calculate a SHA256 hash for a Block.
func (b *Block) CalculateHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// New Block generates a new Block
func NewBlock(prevBlockHash []byte, data string) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.CalculateHash()
	return block
}
