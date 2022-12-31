package main

import "fmt"

func ExchangeCoin(amount int) []int {
	if amount == 0 {
		return []int{}
	}
	var ans []int
	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	for amount > 0 {
		for i := 0; i < len(coins); i++ {
			if amount >= coins[i] {
				ans = append(ans, coins[i])
				amount -= coins[i]
				break
			}
		}
	}
	return ans // TODO: replace this
}

func main() {
	coin := 5000
	fmt.Println(ExchangeCoin(coin))
}