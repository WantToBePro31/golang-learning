package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	counter := 0
	for _, val := range text {
		if val == 'R' || val == 'r' || val == 'S' || val == 's' || val == 'T' || val == 't' || val == 'Z' || val == 'z'{
			counter++
		}
	}
	return counter // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
