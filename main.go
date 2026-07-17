package main

import(
	"fmt"
	"toy-blockchain/cli"

	"toy-blockchain/storage"
)

func main() {

	fmt.Println()
	fmt.Println("Toy Blockchain Started.....")


	// Print Menu
	
	bc := storage.InitializeBlockchain()
	cli.Start(bc)
	
}