package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	newStr := ""
	for ind, val := range data {
		if strings.Contains(val, input) {
			newStr += val
			if ind != len(data)-1 {
				newStr += ","
			}
		}
	}
	if newStr[len(newStr)-1] == ',' {
		newStr = newStr[:len(newStr)-1]
	}
	return newStr // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
