package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	var monthStr string
	switch month {
		case 1:
			monthStr = "January"
		case 2:
			monthStr = "February"
		case 3:
			monthStr = "March"
		case 4:
			monthStr = "April"
		case 5:
			monthStr = "May"
		case 6:
			monthStr = "June"
		case 7:
			monthStr = "July"
		case 8:
			monthStr = "August"
		case 9:
			monthStr = "September"
		case 10:
			monthStr = "October"
		case 11:
			monthStr = "November"
		default:
			monthStr = "December"
	}
	if day < 10 {
		return fmt.Sprintf("0%d-%s-%d", day, monthStr, year)
	}
	return fmt.Sprintf("%d-%s-%d", day, monthStr, year) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
