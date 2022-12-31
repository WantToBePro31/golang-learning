package main

import (
	"fmt"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vokal, konsonan := 0, 0
	for _, val := range str {
		if (val >= 65 && val <= 90) || (val >= 97 && val <= 122) {
			if val == 'A' || val == 'a' || val == 'I' || val == 'i' || val == 'U' || val == 'u' || val == 'E' || val == 'e' || val == 'O' || val == 'o' {
				vokal++
			} else {
				konsonan++
			}
		}
	}
	if vokal == 0 || konsonan == 0 {
		return vokal, konsonan, true
	}
	return vokal, konsonan, false // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
