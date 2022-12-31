package main

import "fmt"

func TicketPlayground(height, age int) int {
	if age > 12 {
		return 100000
	}
	if (age == 12) || (height > 160) {
		return 60000
	}
	if (age >= 10 && age <= 11) || (height > 150) {
		return 40000
	}
	if (age >= 8 && age <= 9) || (height > 135) {
		return 25000
	}
	if (age >= 5 && age <= 7) || (height > 120) {
		return 15000
	}
	return -1 // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
