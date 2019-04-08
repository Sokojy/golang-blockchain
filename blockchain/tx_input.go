package blockchain

import (
	"bytes"
	"golang-blockchain/wallet"
)

// TXInput represents a transaction input
type TXInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

// UsesKey checks whether the address initiated the transaction
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
