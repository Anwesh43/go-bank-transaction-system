package main

import (
	"fmt"

	"bankapp/models"
	"bankapp/utils"
)

func main() {
	accountMap := make(map[string]*models.Account)
	b := models.Bank{
		Name:       "MyBank",
		AccountMap: accountMap,
	}
	b.Greet()

	b.AddAccount("A1", 1000)
	b.AddAccount("A2", 2000)
	b.AddAccount("A3", 3000)

	transactions := make([]models.Transaction, 0)

	transactions = utils.AppendTransaction(transactions, "A1", "deposit", 500)
	transactions = utils.AppendTransaction(transactions, "A2", "withdraw", 300)
	transactions = utils.AppendTransaction(transactions, "A3", "withdraw", 300)
	transactions = utils.AppendTransaction(transactions, "A2", "deposit", 200)
	transactions = utils.AppendTransaction(transactions, "A1", "withdraw", 200)
	transactions = utils.AppendTransaction(transactions, "A3", "withdraw", 500)

	ch := make(chan bool)

	go utils.Transaction(b, transactions, ch)
	<-ch
	fmt.Println("Final Amounts")
	b.GetAccount("A1").PrintDetails()
	b.GetAccount("A2").PrintDetails()
	b.GetAccount("A3").PrintDetails()
}
