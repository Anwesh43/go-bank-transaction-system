package main

import (
	"fmt"
	"time"
)

type Account struct {
	amount int64
	name   string
}

func (a *Account) deposit(amount int64) {
	a.amount += amount
}

func (a *Account) withDraw(amount int64) {
	a.amount -= amount
}

func (a *Account) isAmountAvailable(amount int64) bool {
	return a.amount > amount
}

func (a *Account) printDetails() {
	fmt.Println("Name ", a.name, "Final Amount", a.amount)
}

type Bank struct {
	name       string
	accountMap map[string]*Account
}

func (b *Bank) addAccount(name string, money int64) {
	account := &Account{
		amount: money,
		name:   name,
	}

	b.accountMap[name] = account
}

func (b *Bank) greet() {
	fmt.Println("Welcome to", b.name)
}

func (b *Bank) getAccount(name string) *Account {
	account := b.accountMap[name]
	return account
}

func transaction(b Bank, transactions []Transaction, c chan bool) {
	for _, transaction := range transactions {
		account := b.getAccount(transaction.account_name)
		amount := transaction.amount
		if transaction.transaction_type == "deposit" {
			account.deposit(amount)
		} else {
			if account.isAmountAvailable(amount) {
				account.withDraw(amount)
			} else {
				fmt.Println("amount is not available for", transaction.account_name)
			}
		}
		fmt.Println("Ran 1 transaction for", transaction.account_name)
		account.printDetails()
		time.Sleep(time.Second)
	}
	c <- true
}

type Transaction struct {
	transaction_type string
	account_name     string
	amount           int64
}

func appendTransaction(transactions []Transaction, account_name string, transaction_type string, amount int64) []Transaction {
	return append(transactions, Transaction{
		account_name:     account_name,
		transaction_type: transaction_type,
		amount:           amount,
	})
}

func main() {
	accountMap := make(map[string]*Account)
	b := Bank{
		name:       "MyBank",
		accountMap: accountMap,
	}
	b.greet()

	b.addAccount("A1", 1000)
	b.addAccount("A2", 2000)
	b.addAccount("A3", 3000)

	transactions := make([]Transaction, 0)

	transactions = appendTransaction(transactions, "A1", "deposit", 500)
	transactions = appendTransaction(transactions, "A2", "withdraw", 300)
	transactions = appendTransaction(transactions, "A3", "withdraw", 300)
	transactions = appendTransaction(transactions, "A2", "deposit", 200)
	transactions = appendTransaction(transactions, "A1", "withdraw", 200)
	transactions = appendTransaction(transactions, "A3", "withdraw", 500)

	ch := make(chan bool)

	go transaction(b, transactions, ch)
	<-ch
	fmt.Println("Final Amounts")
	b.getAccount("A1").printDetails()
	b.getAccount("A2").printDetails()
	b.getAccount("A3").printDetails()
}
