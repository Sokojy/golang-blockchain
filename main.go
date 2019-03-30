package main

import (
	"fmt"
	"golang-blockchain/blockchain"
	"strconv"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.AddBlock("First block")
	bc.AddBlock("Second block")
	bc.AddBlock("Third block")

	for _, block := range bc.Blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Prev. hash: %x\n", block.HashPrevBlock)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("--------------------------------------------------------")
	}
}
