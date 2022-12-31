package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	average := float64(math+science+english+indonesia)/4.0
	if average == 100.0 {
		return "Sempurna"
	}
	if average >= 90 {
		return "Sangat Baik"
	}
	if average >= 80 {
		return "Baik"
	}
	if average >= 70 {
		return "Cukup"
	}
	if average >= 60 {
		return "Kurang"	
	}
	
	return "Sangat kurang" // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
