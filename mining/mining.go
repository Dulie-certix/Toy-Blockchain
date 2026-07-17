package mining

import (
	"fmt"
	"strings"
	"toy-blockchain/block"
)

func Mine(b *block.Block, difficulty int) {

	target := strings.Repeat("0", difficulty)

	_= target

	fmt.Println("Mining Started........")
	for {
		b.Hash = b.CalculateHash()

		if strings.HasPrefix(b.Hash,target){
			fmt.Println("Mining Finished!")
			fmt.Println("Nonce:", b.Nonce)
			fmt.Println("Hash :", b.Hash)
			break
			
		}

		b.Nonce++
	}
}
