package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {
	tmp := []rune(str)
	sz := len(str)
	for i := 0; i < sz/2; i++ {
		tmp[i], tmp[sz-i-1] = tmp[sz-i-1], tmp[i]
	}
	return string(tmp) // TODO: replace this
}

func Generate(str string) string {
	str = Reverse(str)
	tmp := []rune(str)
	for ind, val := range tmp {
		if val == 'a' || val == 'A' || val == 'e' || val == 'E' {
			tmp[ind] += 4
		} else if val == 'i' || val == 'I' || val == 'o' || val == 'O' {
			tmp[ind] += 6
		} else if val == 'u' || val == 'U' {
			tmp[ind] -= 20
		}
		if val >= 65 && val <= 90 {
			tmp[ind] += 32
		} else if val >= 97 && val <= 122 {
			tmp[ind] -= 32
		}
	}
	return strings.ReplaceAll(string(tmp), " ", "") // TODO: replace this
}

func CheckPassword(str string) string {
	if len(str) < 7 {
		return "sangat lemah"
	}
	symbol := 0
	for _, val := range str {
		if (val < 65 || val > 90) && (val < 97 || val > 122) && (val < 48 || val > 57) {
			symbol++
		}
	}
	if symbol == 0 {
		return "lemah"
	}
	if len(str) >= 14 {
		return "kuat"
	}
	return "sedang" // TODO: replace this
}

func PasswordGenerator(base string) (string, string) {
	generatedPassword := Generate(base)
	return generatedPassword, CheckPassword(generatedPassword) // TODO: replace this
}

func main() {
	data := "Hello World" // bisa digunakan untuk melakukan debug

	res, check := PasswordGenerator(data)

	fmt.Println(res)
	fmt.Println(check)
}
