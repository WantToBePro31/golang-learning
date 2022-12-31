package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)


func AdvanceCalculator(calculate string) float32 {
	calculateStr := strings.Split(calculate, " ")
	baseFloat, _ := strconv.ParseFloat(calculateStr[0], 64)
	c := internal.NewCalculator(float32(baseFloat))
	operator := ""
	for i := 1; i < len(calculateStr); i++ {
		if i % 2 == 1 {
			operator = calculateStr[i]
		} else {
			num, _ := strconv.ParseFloat(calculateStr[i], 64)
			if operator == "+" {
				c.Add(float32(num))
			} else if operator == "-" {
				c.Subtract(float32(num))
			} else if operator == "*" {
				c.Multiply(float32(num))
			} else {
				c.Divide(float32(num))
			}
		}
	}
	return c.Result() // TODO: replace this
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
