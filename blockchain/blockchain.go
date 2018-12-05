package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.GetLastBlock()
	newBlock := NewBlock(prevBlock.Hash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// First Block in the chain is called Genesis Block
func NewGenesisBlock() *Block {
	return NewBlock([]byte{}, "Genesis Block")
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
