package cli

import (
	"fmt"
	"toy-blockchain/blockchain"
	"toy-blockchain/ledger"
	"toy-blockchain/validation"
	"toy-blockchain/storage"
	"toy-blockchain/transaction"
) 

func Start(bc *blockchain.Blockchain) {

	for {
		ShowMenu()

		choice := ReadChoice()
		fmt.Println("============================")

		switch choice {

		case 1:
			fmt.Println("\nAdd Transaction Selected")

			tx := ReadTransaction()
			bc.AddBlock([]transaction.Transaction{tx})

			fmt.Println("Transaction added successfully")

		case 2:
			fmt.Println("\nPrint Ledger Selected")
			ledger.PrintLedger(bc)
			fmt.Println("Print Ledger successfully")


		case 3:
			fmt.Println("\nValidate Blockchain Selected")
			if validation.IsValid(bc){
				fmt.Println("Blockchain is valid")
			} else {
				fmt.Println("Blockchain is invalid")
			}
			

		case 4:
			fmt.Println("\nSave Blockchain Selected")
			storage.SaveBlockchain(bc)
			fmt.Println("Blockchain saved successfully")


		case 5:
			fmt.Println("\nExiting...")
			fmt.Println("\nThank you for using Toy Blockchain!")
			return

		default:
			fmt.Println("\nInvalid Choice. Please try again.")
		
		}
		fmt.Println()
	}

}