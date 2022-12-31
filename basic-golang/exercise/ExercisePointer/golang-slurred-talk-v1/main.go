package main

import "fmt"

func SlurredTalk(words *string) {
	// TODO: answer here
	runeWords := []rune(*words)
	for i := 0; i < len(runeWords); i++ {
		char := runeWords[i]
		if char == 'S' ||  char == 'R' || char == 'Z' {
			runeWords[i] = 'L'
		} else if char == 's' || char == 'r' || char == 'z' {
			runeWords[i] = 'l'
		}
	}
	*words = string(runeWords)
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Steven"
	SlurredTalk(&words)
	fmt.Println(words)
}
