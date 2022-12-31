package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	if gender == "laki-laki" {
		return (float64(height) - 100.0) - ((float64(height) - 100.0) * 0.1)
	}
	return (float64(height) - 100.0) - ((float64(height) - 100.0) * 0.15) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
