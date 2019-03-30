package blockchain

// Blockchain represents a blockchain
type Blockchain struct {
	Blocks []*Block
}

// AddBlock is used to add a new block to blockchain
func (b *Blockchain) AddBlock(data string) {
	hashPrevBlock := b.Blocks[len(b.Blocks)-1].Hash
	block := NewBlock(data, hashPrevBlock)
	b.Blocks = append(b.Blocks, block)
}

// NewBlockchain creates a new blockchain with genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis block", []byte{})
	blockchain := &Blockchain{[]*Block{genesisBlock}}
	return blockchain
}
