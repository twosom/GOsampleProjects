package main

import (
	"fmt"
	"github.com/twosom/sampleProjects/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("twosom")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Balance())
}
