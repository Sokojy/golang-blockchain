package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block represents a block
type Block struct {
	Time          int64
	Hash          []byte
	HashPrevBlock []byte
	Data          []byte
}

// NewBlock is used to crerate a new block
func NewBlock(data string, hashPrevBlock []byte) *Block {
	time := time.Now().Unix()
	hash := CreateHash(data, hashPrevBlock, time)
	block := &Block{
		Data:          []byte(data),
		HashPrevBlock: hashPrevBlock,
		Hash:          hash[:],
		Time:          time,
	}
	return block
}

// CreateHash will calculate and return hash
func CreateHash(data string, hashPrevBlock []byte, time int64) []byte {
	timestamp := []byte(strconv.FormatInt(time, 10))
	headers := bytes.Join([][]byte{hashPrevBlock, []byte(data), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}

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

func main() {
	blockchain := NewBlockchain()
	blockchain.AddBlock("First block")
	blockchain.AddBlock("Second block")
	blockchain.AddBlock("Third block")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Prev. hash: %x\n", block.HashPrevBlock)
		fmt.Println("--------------------------------------------------------")
	}
}
