package block

import (
	"fmt"
	"time"
	"toy-blockchain/transaction"
	"toy-blockchain/utils"
)

type Block struct {
	Index		int
	Timestamp   int64
	Transactions []transaction.Transaction
	PreviousHash  string
	Nonce       int
	Hash		string
}

// Create a Root of the block
func NewGenesisBlock() Block{
	block := Block{
		Index: 0,
		Timestamp: time.Now().Unix(),
		Transactions: []transaction.Transaction{},
		PreviousHash: "00000000000000000000000000000000",
		Nonce: 0,
	}

	block.Hash = block.CalculateHash()

	return block
}

// function for struct
func(b*Block) CalculateHash() string { 

	data := fmt.Sprintf(
		"%d%d%v%s%d", 
		b.Index, 
		b.Timestamp, 
		b.Transactions,
		b.PreviousHash, 
		b.Nonce,
	)

	return utils.CalculateHash(data)
}