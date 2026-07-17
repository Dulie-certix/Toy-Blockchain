package validation

import (
	"toy-blockchain/blockchain"
)

func IsValid(bc *blockchain.Blockchain) bool{

	for i :=1; i<len(bc.Blocks); i++{

		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		if current.Hash != current.CalculateHash(){
			return false
		}

		if current.PreviousHash != previous.Hash{
			return false
		}
	
	}

	return true
}