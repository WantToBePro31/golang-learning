package main

import (
	"fmt"
	"strconv"
)

func ChangeToCurrency(change int) string {
	strChange := strconv.Itoa(change)
	sz := len(strChange)
	if sz % 3 == 0 {
		sz += sz/3 - 1
	} else {
		sz += sz/3
	}
	rupiahs := make([]byte, sz)
	triple := 0
	sz--
	for i := len(strChange)-1; i >= 0; i-- {
		rupiahs[sz] = strChange[i]
		sz--
		triple++
		if triple == 3 && sz != -1{
			rupiahs[sz] = '.'
			sz--
			triple = 0
		}
	}
	return fmt.Sprintf("Rp. %s", string(rupiahs)) // TODO: replace this
}

func MoneyChange(money int, listPrice ...int) string {
	totalPrice := 0
	for _, val := range listPrice {
		totalPrice += val
	}
	if money < totalPrice {
		return "Uang tidak cukup"
	}
	return ChangeToCurrency(money - totalPrice) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(MoneyChange(100000, 50000, 10000, 10000, 5000, 5000))
}