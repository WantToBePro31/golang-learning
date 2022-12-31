package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {
	// TODO: answer here
	if number[0] == '6' && number[1] == '2' && number[2] == '8' {
		if len(number) >= 11 {
			if number[3] == '1' {
				if number[4] >= '1' && number[4] <= '5' {
					*result = "Telkomsel"
				} else if number[4] >= '6' && number[4] <= '9' {
					*result = "Indosat"
				}
			} else if number[3] == '2' {
				if number[4] >= '1' && number[4] <= '3' {
					*result = "XL"
				} else if number[4] >= '7' && number[4] <= '9' {
					*result = "Tri"
				}
			} else if (number[3] == '5' && number[4] == '2') || (number[3] == '5' && number[4] == '3') {
				*result = "AS"
			} else if number[3] == '8' {
				if number[4] >= '1' && number[4] <= '8' {
					*result = "Smartfren"
				}
			} else {
				*result = "invalid"
			}
		} else {
			*result = "invalid"
		}
	} else if number[0] == '0' && number[1] == '8' {
		if len(number) >= 10 {
			if number[2] == '1' {
				if number[3] >= '1' && number[3] <= '5' {
					*result = "Telkomsel"
				} else if number[3] >= '6' && number[4] <= '9' {
					*result = "Indosat"
				}
			} else if number[2] == '2' {
				if number[3] >= '1' && number[3] <= '3' {
					*result = "XL"
				} else if number[3] >= '7' && number[3] <= '9' {
					*result = "Tri"
				}
			} else if (number[2] == '5' && number[3] == '2') || (number[2] == '5' && number[3] == '3') {
				*result = "AS"
			} else if number[2] == '8' {
				if number[3] >= '1' && number[3] <= '8' {
					*result = "Smartfren"
				}
			} else {
				*result = "invalid"
			}
		} else {
			*result = "invalid"
		}
	} else {
		*result = "invalid"
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "081211111111"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
