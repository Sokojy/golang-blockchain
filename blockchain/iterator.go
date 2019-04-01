package blockchain

import (
	"log"

	"github.com/dgraph-io/badger"
)

// Iterator represents a blockchain iterator
type Iterator struct {
	currentHash []byte
	DB          *badger.DB
}

// Next returns a next block starting from tip
func (iter *Iterator) Next() *Block {
	var block *Block

	err := iter.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.currentHash)
		if err != nil {
			return err
		}
		var encodedBlock []byte
		item.Value(func(val []byte) error {
			encodedBlock = val
			return nil
		})
		block = Deserialize(encodedBlock)

		return err
	})
	if err != nil {
		log.Panic(err)
	}

	iter.currentHash = block.HashPrevBlock

	return block
}
