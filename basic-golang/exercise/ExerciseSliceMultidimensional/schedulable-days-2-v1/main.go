package main

import (
	"fmt"
)

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	} 
	sz := len(villager)
	date := make(map[int]int)
	for i := 0; i < len(villager); i++ {
		for j := 0; j < len(villager[i]); j++ {
			date[villager[i][j]]++
		}
	}
	var ans []int
	for val, counter := range date {
		if counter == sz {
			ans = append(ans, val)
		}
	}
	return ans // TODO: replace this
}

func main() {
	datas := [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}
	fmt.Println(SchedulableDays(datas))
}
