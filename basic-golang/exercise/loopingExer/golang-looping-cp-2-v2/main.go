package main

import (
	"fmt"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	newStr := ""
	for i := len(str)-1; i >= 1; i-- {
		newStr += string(str[i])
		if str[i-1] != ' ' && str[i] != ' ' {
			newStr += "_"
		}
	}
	newStr += string(str[0])
	return newStr // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
