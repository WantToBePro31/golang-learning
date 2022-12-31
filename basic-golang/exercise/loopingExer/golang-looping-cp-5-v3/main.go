package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	if str == "" {
		return str
	}
	newStr := ""
	pivot := 0
	for last, val := range str {
		if val == ' ' {
			if last-pivot == 1 {
				newStr += string(str[pivot])
			} else {
				str1, str2 := string(str[pivot]), string(str[last-1])
				if unicode.IsUpper(rune(str[pivot])) {
					str1 = strings.ToLower(str1)
					str2 = strings.ToUpper(str2)
				}
				newStr += str2
				for i := last-2; i >= pivot+1; i-- {
					newStr += strings.ToLower(string(str[i]))
				}
				newStr += str1
			}
			newStr += " "
			pivot = last+1
		}
	}
	if len(str)-1-pivot == 1 {
		newStr += string(str[pivot])
	} else {
		str1, str2 := string(str[pivot]), string(str[len(str)-1])
		if unicode.IsUpper(rune(str[pivot])) {
			str1 = strings.ToLower(str1)
			str2 = strings.ToUpper(str2)
		}
		newStr += str2
		for i := len(str)-2; i >= pivot+1; i-- {
			newStr += strings.ToLower(string(str[i]))
		}
		newStr += str1
	}
	return newStr // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
}
