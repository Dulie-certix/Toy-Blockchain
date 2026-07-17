package storage

import (
	"encoding/json"
	"os"

	"toy-blockchain/block"
	"toy-blockchain/blockchain"
)

func LoadBlockchain() (*blockchain.Blockchain, error){

	data, err := os.ReadFile("ledger.json")

	if err != nil{
		return nil, err
	}

	var blocks []block.Block

	err = json.Unmarshal(data, &blocks)

	if err != nil{
		return nil, err
	}

	bc := &blockchain.Blockchain{
		Blocks: blocks,
	}
	return bc, nil

}

func InitializeBlockchain() *blockchain.Blockchain{
	
	bc, err := LoadBlockchain()
	
	if err != nil{
		return blockchain.NewBlockchain()
	}
	return bc
}