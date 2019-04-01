package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Block represents a block
type Block struct {
	Time          int64
	Hash          []byte
	HashPrevBlock []byte
	Data          []byte
	Nonce         int
}

// NewBlock is used to crerate a new block
func NewBlock(data string, hashPrevBlock []byte) *Block {
	time := time.Now().Unix()
	block := &Block{
		Data:          []byte(data),
		HashPrevBlock: hashPrevBlock,
		Hash:          []byte{},
		Time:          time,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

// Serialize serializes a block
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return res.Bytes()
}

// Deserialize deserializes a block
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
