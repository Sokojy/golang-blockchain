package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// Transaction represents a transaction
type Transaction struct {
	ID      []byte
	Outputs []TXOutput
	Inputs  []TXInput
}

// TXOutput represents a transaction output
type TXOutput struct {
	Value  int
	PubKey string
}

// TXInput represents a transaction input
type TXInput struct {
	ID  []byte
	Out int
	Sig string
}

// NewTransaction creates a new transaction
func NewTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	accumulated, validOutputs := bc.FindSpendableOutputs(from, amount)

	if accumulated < amount {
		log.Panic("Not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TXOutput{amount, to})
	if accumulated > amount {
		outputs = append(outputs, TXOutput{accumulated - amount, from})
	}

	tx := Transaction{nil, outputs, inputs}
	tx.SetID()

	return &tx
}

// CoinbaseTX creates a new coinbase transaction
func CoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coin to %s", to)
	}
	txinput := TXInput{[]byte{}, -1, data}
	txoutput := TXOutput{100, to}
	transaction := Transaction{nil, []TXOutput{txoutput}, []TXInput{txinput}}

	return &transaction
}

// SetID sets id to transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	if err != nil {
		log.Panic()
	}

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// CanUnlockOutputWith checks whether the address initiated the transaction
func (in *TXInput) CanUnlockOutputWith(data string) bool {
	return in.Sig == data
}

// CanBeUnlockedWith checks if the output can be unlocked with the provided data
func (out *TXOutput) CanBeUnlockedWith(data string) bool {
	return out.PubKey == data
}

// IsCoinbase is used to chech transaction
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}
