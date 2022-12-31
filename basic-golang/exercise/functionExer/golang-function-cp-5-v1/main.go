package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	min := 100001
	for _, val := range nums {
		if val < min {
			min = val
		}
	}
	return min // TODO: replace this
}

func FindMax(nums ...int) int {
	max := -1
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	return max // TODO: replace this
}

func SumMinMax(nums ...int) int {
	return FindMin(nums...) + FindMax(nums...) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
