package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	if len(data) == 0 {
		return []int{}
	}
	ans := make([]int, len(data[0]))
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			ans[j] += data[i][j][0] - data[i][j][1]
		}
	}
	return ans // TODO: replace this
}

func main() {
	datas := [][][2]int{{{1000, 500}, {500, 200}}, {{1200, 200}, {1000, 800}}, {{500, 100}, {700, 100}}}
	fmt.Println(CountProfit(datas))
}
