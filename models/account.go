package models

import "fmt"

type Account struct {
	Amount int64
	Name   string
}

func (a *Account) Deposit(amount int64) {
	a.Amount += amount
}

func (a *Account) WithDraw(amount int64) {
	a.Amount -= amount
}

func (a *Account) IsAmountAvailable(amount int64) bool {
	return a.Amount > amount
}

func (a *Account) PrintDetails() {
	fmt.Println("Name ", a.Name, "Final Amount", a.Amount)
}
