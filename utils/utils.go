package utils

import (
	"bankapp/models"
	"fmt"
	"time"
)

func Transaction(b models.Bank, transactions []models.Transaction, c chan bool) {
	for _, transaction := range transactions {
		account := b.GetAccount(transaction.Account_name)
		amount := transaction.Amount
		if transaction.Transaction_type == "deposit" {
			account.Deposit(amount)
		} else {
			if account.IsAmountAvailable(amount) {
				account.WithDraw(amount)
			} else {
				fmt.Println("amount is not available for", transaction.Account_name)
			}
		}
		fmt.Println("Ran 1 transaction for", transaction.Account_name)
		account.PrintDetails()
		time.Sleep(time.Second)
	}
	c <- true
}

func AppendTransaction(transactions []models.Transaction, account_name string, transaction_type string, amount int64) []models.Transaction {
	return append(transactions, models.Transaction{
		Account_name:     account_name,
		Transaction_type: transaction_type,
		Amount:           amount,
	})
}
