package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	if len(products) == 0 {	
		return []int{}
	}
	totalPrice := 0
	for _, prod := range products {
		totalPrice += prod.Price + prod.Tax
	}
	if amount == totalPrice {
		return []int{}
	}
	amount -= totalPrice
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
	money := 10000
	products := []Product{{Name: "Baju", Price: 5000, Tax: 500}, {Name: "Celana", Price: 3000, Tax: 300}}
	fmt.Println(MoneyChanges(money, products))
}