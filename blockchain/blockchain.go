package blockchain

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

// Blockchain represents a blockchain
type Blockchain struct {
	LastHash []byte
	DB       *badger.DB
}

// InitBlockchain is used to init blockchain
func InitBlockchain() *Blockchain {
	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := NewBlock("Genesis", []byte{})
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		}
		item, err := txn.Get([]byte("lh"))
		if err != nil {
			return err
		}
		item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		return err
	})

	if err != nil {
		log.Panic(err)
	}

	blockchain := Blockchain{lastHash, db}
	return &blockchain
}

// AddBlock adds a new block to blockchain
func (chain *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := chain.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		if err != nil {
			return err
		}
		item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})

		return err
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = chain.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	if err != nil {
		log.Panic(err)
	}
}

// Iterator creates a new blockchain iterator
func (chain *Blockchain) Iterator() *Iterator {
	iter := &Iterator{chain.LastHash, chain.DB}

	return iter
}
