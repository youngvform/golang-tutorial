package account

import "log"

type Account struct {
	owner string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) Deopsit(amount int)  {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) {
	if a.balance < amount {
		log.Fatalln("can't withdraw")
	}
	a.balance -= amount
}

func (a Account) Balance() int {
	return a.balance
}