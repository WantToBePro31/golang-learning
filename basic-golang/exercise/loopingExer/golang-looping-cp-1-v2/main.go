package main

import "fmt"

func CountingNumber(n int) float64 {
	counter, add, increment := 0.0, 1.0, 0.5
	for add <= float64(n) {
		counter += add
		add += increment
	}
	return counter // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
