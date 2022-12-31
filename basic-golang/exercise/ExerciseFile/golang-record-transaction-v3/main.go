package main

import (
	"fmt"
	"os"
	"sort"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	if len(transactions) == 0 {
		return nil
	}
	file, errOpen := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if errOpen != nil {
		return errOpen
	}
	defer file.Close()
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date < transactions[j].Date
	})
	income, expense, date := 0, 0, transactions[0].Date
	for _, val := range transactions {
		if val.Date == date {
			if val.Type == "income" {
				income += val.Amount
			} else {
				expense += val.Amount
			}
		} else {
			var text string
			if income >= expense {
				text = fmt.Sprintf("%s;income;%d\n", date, income-expense)
			} else {
				text = fmt.Sprintf("%s;expense;%d\n", date, expense-income)
			}
			_, errWrite := file.WriteString(text)
			if errWrite != nil {
				return errWrite
			}
			if val.Type == "income" {
				income, expense = val.Amount, 0
			} else {
				income, expense = 0, val.Amount
			}
			date = val.Date
		}
	}
	var text string
	if income >= expense {
		text = fmt.Sprintf("%s;income;%d", date, income-expense)
	} else {
		text = fmt.Sprintf("%s;expense;%d", date, expense-income)
	}
	_, errWrite := file.WriteString(text)
	if errWrite != nil {
		return errWrite
	}
	return nil // TODO: replace this
}


func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
