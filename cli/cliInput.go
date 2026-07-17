package cli

import (
	"fmt"
	"toy-blockchain/transaction"
)

func ReadChoice() int{

	var choice int


	_, err := fmt.Scan(&choice)

	if err != nil{
		return -1
	}

	return choice
}

func ReadTransaction() transaction.Transaction{

	var sender,recipient string
	var amount float64

	fmt.Print("Sender: ")
	fmt.Scan(&sender)

	fmt.Print("Recipient: ")
	fmt.Scan(&recipient)

	fmt.Print("Amount: ")
	fmt.Scan(&amount)

	return transaction.NewTransaction(sender, recipient, amount)
}