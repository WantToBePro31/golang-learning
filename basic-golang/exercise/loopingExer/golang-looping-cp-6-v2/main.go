package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	max, ans := -1, 0
	strNumber := strconv.Itoa(numbers)
	for i := 0; i < len(strNumber)-1; i++ {
		bil1, bil2 := int(strNumber[i]) - 48, int(strNumber[i+1]) - 48
		if bil1 + bil2 > max {
			max = bil1 + bil2
			ans = bil1*10 + bil2 
		}
	}
	return ans // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
