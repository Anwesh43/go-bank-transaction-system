package models

import "fmt"

type Bank struct {
	Name       string
	AccountMap map[string]*Account
}

func (b *Bank) AddAccount(name string, money int64) {
	account := &Account{
		Amount: money,
		Name:   name,
	}

	b.AccountMap[name] = account
}

func (b *Bank) Greet() {
	fmt.Println("Welcome to", b.Name)
}

func (b *Bank) GetAccount(name string) *Account {
	account := b.AccountMap[name]
	return account
}
