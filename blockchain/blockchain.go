package blockchain

import (
	"time"
	"toy-blockchain/block"
	"toy-blockchain/transaction"
	"toy-blockchain/mining"
)

type Blockchain struct {
	Blocks []block.Block
}

// Create a Blockchain
func NewBlockchain() *Blockchain {

	genesis := block.NewGenesisBlock()

	return &Blockchain{
		Blocks: []block.Block{genesis},
	}
}  

//Get leatst block Index
func (bc *Blockchain) GetLatestBlock() block.Block{
	return bc.Blocks[len(bc.Blocks)-1]
}

// Add a new block
func (bc *Blockchain) AddBlock(transactions []transaction.Transaction) {

	latest := bc.GetLatestBlock()
	
	newBlock := block.Block{
		Index: latest.Index + 1,
		Timestamp: time.Now().Unix(),
		Transactions: transactions,
		PreviousHash: latest.Hash,
		Nonce: 0,
	}

	// Mining
	mining.Mine(&newBlock, 3)

	// Save
	bc.Blocks = append(bc.Blocks, newBlock)
}
