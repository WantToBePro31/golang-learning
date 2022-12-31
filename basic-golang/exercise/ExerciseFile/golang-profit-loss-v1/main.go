package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	var datas []string
	fullData, err := os.ReadFile(path)
	if err != nil {
		return []string{}, err
	} 
	if string(fullData) == "" {
		return []string{}, nil
	}
	everyData := strings.Split(string(fullData), "\n")
	for _, val := range everyData {
		datas = append(datas, val)
	}
	return datas, nil // TODO: replace this
}

func CalculateProfitLoss(data []string) string {
	income, expense, date := 0, 0, ""
	for _, val := range data {
		formatStr := strings.Split(val, ";")
		date = formatStr[0]
		amount, _ := strconv.Atoi(formatStr[2])
		if formatStr[1] == "income" {
			income += amount
		} else {
			expense += amount
		}
	}
	if income >= expense {
		return fmt.Sprintf("%s;profit;%d", date, income-expense)
	}
	return fmt.Sprintf("%s;loss;%d", date, expense-income) // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
