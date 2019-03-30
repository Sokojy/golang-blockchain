package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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

// CreateHash will calculate and return hash
func CreateHash(data string, hashPrevBlock []byte, time int64) []byte {
	timestamp := []byte(strconv.FormatInt(time, 10))
	headers := bytes.Join([][]byte{hashPrevBlock, []byte(data), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}
