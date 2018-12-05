package blockchain


type Blockchain struct {
	blocks []*Block
}


func (bc *Blockchain) GetLastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.GetLastBlock()
	newBlock := NewBlock(prevBlock.Hash, data)
	bc.blocks = append(bc.blocks, newBlock)
}

// First Block in the chain is called Genesis Block
func NewGenesisBlock() *Block {
	return NewBlock([]byte{}, "Genesis Block")
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
