package ledger

import (
	"fmt"
	"toy-blockchain/blockchain"
)

func PrintLedger(bc *blockchain.Blockchain){

	fmt.Println("\n========== BLOCKCHAIN LEDGER ==========")

	for _, block := range bc.Blocks{

		fmt.Println("--------------------------------------")
		fmt.Println("Block Index :", block.Index)
		fmt.Println("Timestamp   :", block.Timestamp)
		fmt.Println("PreviousHash:", block.PreviousHash)
		fmt.Println("Hash        :", block.Hash)
		fmt.Println("Nonce       :", block.Nonce)
	

		fmt.Println("Transactions:")

		if len(block.Transactions) == 0{
			fmt.Println("No Transactions")
		}else{
			
			for _, tx := range block.Transactions{
				
				fmt.Printf(" %s -> %s : %.2f\n",
					tx.Sender,
					tx.Recipient,
					tx.Amount )
			}
		}
	}

	fmt.Println("==========================================")
}