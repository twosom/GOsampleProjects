package main

import (
	"fmt"
	"github.com/twosom/sampleProjects/accounts"
)

func main() {
	account := accounts.NewAccount("twosom")
	fmt.Println(account)
}
