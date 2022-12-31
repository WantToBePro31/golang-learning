package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

type Borrower struct {
	Id string `json:"id"`
	TotalLoan int `json:"total_loan"`
}

// json structure
type LoanRecord struct {
	MonthDate string `json:"month_date"`
	StartBalance int `json:"start_balance"`
	EndBalance int `json:"end_balance"`
	Borrowers []Borrower `json:"borrowers"`
}

func LoanReport(data LoanData) LoanRecord {
	dates := strings.Split(data.Data[0].Date, "-")
	month := fmt.Sprintf("%s %s", dates[1], dates[2])
	employees := make(map[string]int)
	counter := 0
	for _, val := range data.Data {
		for _, id := range val.EmployeeIDs {
			employees[id]++
			counter++
			if counter * 50000 >= data.StartBalance {
				break
			}
		}
	}
	loan := 0
	var borrowers []Borrower
	for id, countLoan := range employees {
		var borrower Borrower
		borrower.Id = id
		borrower.TotalLoan = countLoan * 50000
		loan += borrower.TotalLoan
		borrowers = append(borrowers, borrower)
	}
	sort.Slice(borrowers, func(i, j int) bool {
		idA, _ :=  strconv.Atoi(borrowers[i].Id)
		idB, _ :=  strconv.Atoi(borrowers[j].Id)
		return idA < idB
	})
	endBalance := data.StartBalance - loan
	return LoanRecord{MonthDate: month, StartBalance: data.StartBalance, EndBalance: endBalance, Borrowers: borrowers} // TODO: replace this
}

func RecordJSON(record LoanRecord, path string) error {
	jsonData, err := json.Marshal(record)
	if err != nil {
		return err
	}
	
	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

// gunakan untuk debug
func main() {
	records := LoanReport(LoanData{
		StartBalance: 500000,
		Data: []Loan{
			{"01-January-2021", []string{"1", "2"}},
			{"02-January-2021", []string{"1", "2", "3"}},
			{"03-January-2021", []string{"2", "3"}},
			{"04-January-2021", []string{"1", "3"}},
		},
		Employees: []Employee{
			{"1", "Employee A", "Manager"},
			{"2", "Employee B", "Staff"},
			{"3", "Employee C", "Staff"},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}
