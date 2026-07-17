package storage

import (
	"os"
	"encoding/json"

	"toy-blockchain/blockchain"
)

func SaveBlockchain(bc *blockchain.Blockchain) error{
	
	data, err := json.MarshalIndent(bc.Blocks, "","    ")

	if err != nil{
		return err
	}

	err = os.WriteFile("ledger.json", data, 0644)

	if err != nil {
		return err
	}

	return nil
}