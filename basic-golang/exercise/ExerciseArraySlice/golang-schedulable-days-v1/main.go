package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	var bothDates []int
	for i := 0; i < len(date1); i++ {
		for j := 0; j < len(date2); j++ {
			if date1[i] == date2[j] {
				bothDates = append(bothDates, date1[i])
			}
		}
	}
	if len(bothDates) == 0 {
		return []int{}
	}
	return bothDates // TODO: replace this
}

func main() {
	datesImam := []int{11, 12, 13, 14, 15}
	datesPermana := []int{5, 10, 12, 13, 20, 21}
	fmt.Println(SchedulableDays(datesImam, datesPermana))
}
