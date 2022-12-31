package main

import (
	"fmt"
	"strconv"
)

func GetDayDifference(date string) int {
	month := map[string]int{
		"January": 1,
		"February": 2,
		"March": 3,
		"April": 4,
		"May": 5,
		"June": 6,
		"July": 7,
		"August": 8,
		"September": 9,
		"October": 10,
		"November": 11,
		"December": 12,
	}
	countDate := map[int]int{
		1: 30,
		3: 30,
		4: 30,
		5: 30,
		6: 30,
		7: 30,
		8: 30,
		9: 30,
		10: 30,
		11: 30,
		12: 30,
	}
	var dateDetail []string
	pivot := 0
	for i := 0; i < len(date); i++ {
		if date[i] == ' ' {
			dateDetail = append(dateDetail, date[pivot:i])
			pivot = i+1
		}
	}
	dateDetail = append(dateDetail, date[pivot:])

	year, _ := strconv.Atoi(dateDetail[5])
	if year % 400 == 0 {
		countDate[2] = 29
	} else if year % 100 == 0 {
		countDate[2] = 28
	} else if year % 4 == 0 {
		countDate[2] = 29
	} else {
		countDate[2] = 28
	}

	totalDay, startMonth, endMonth := 0, month[dateDetail[1]], month[dateDetail[4]]
	startDate, _ := strconv.Atoi(dateDetail[0])
	endDate, _ := strconv.Atoi(dateDetail[3])
	if startMonth == endMonth {
		return endDate - startDate + 1
	}
	totalDay += countDate[startMonth] - startDate + 1
	for i := startMonth+1; i <= endMonth-1; i++ {
		totalDay += countDate[i]
	}
	totalDay += endDate
	return totalDay // TODO: replace this
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	countWorker := make(map[string]int)
	for i := 0; i < rangeDay; i++ {
		for j := 0; j < len(data[i]); j++ {
			countWorker[data[i][j]]++
		}
	}
	salaryWorker := make(map[string]string)
	for name, counter := range countWorker {
		salaryWorker[name] = FormatRupiah(counter * 50000)
	}
	return salaryWorker // TODO: replace this
}

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi helper untuk mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	strNumber := strconv.Itoa(number)
	sz := len(strNumber)
	if sz % 3 == 0 {
		sz += sz/3 - 1
	} else {
		sz += sz/3
	}
	rupiahs := make([]byte, sz)
	triple := 0
	sz--
	for i := len(strNumber)-1; i >= 0; i-- {
		rupiahs[sz] = strNumber[i]
		sz--
		triple++
		if triple == 3 && sz != -1{
			rupiahs[sz] = '.'
			sz--
			triple = 0
		}
	}
	return fmt.Sprintf("Rp %s", string(rupiahs)) // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	return GetSalary(GetDayDifference(dateRange), data) // TODO: replace this
}

func main() {
	res := GetSalaryOverview("21 February - 23 February 2021", [][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	fmt.Println(res)
}
