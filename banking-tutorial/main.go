package main

import (
	"fmt"
	"github.com/youngvform/go-turorial/banking-tutorial/account"
)

func main() {
	account := account.NewAccount("kim")
	account.Deopsit(100)
	fmt.Println(account.Balance())
	account.Withdraw(110)
	fmt.Println(account.Balance())

}
