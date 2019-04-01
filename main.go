package main

import (
	"os"

	"golang-blockchain/blockchain"
	"golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockchain()
	defer chain.DB.Close()

	cli := cli.CommandLine{chain}
	cli.Run()
}
