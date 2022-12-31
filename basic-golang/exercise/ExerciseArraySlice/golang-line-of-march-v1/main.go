package main

import "fmt"

func Sortheight(height []int) []int {
	for i := 0; i < len(height); i++ {
		for j := i+1; j < len(height); j++ {
			if height[i] > height[j] {
				height[i], height[j] = height[j], height[i]
			}
		}
	}
	return height // TODO: replace this
}

func main() {
	height := []int{172, 170, 150, 165, 144, 155, 159}
	fmt.Println(Sortheight(height))
}
