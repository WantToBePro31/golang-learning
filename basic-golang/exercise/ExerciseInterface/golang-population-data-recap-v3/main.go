package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	if len(data) == 0 {
		return []map[string]any{}
	}
	var persons []map[string]any
	for _, str := range data {
		personData := strings.Split(str, ";")
		tmp := make(map[string]any)
		tmp["name"] = personData[0]
		age, _ := strconv.Atoi(personData[1])
		tmp["age"] = age
		tmp["address"] = personData[2]
		if personData[3] != "" {
			height, _ := strconv.ParseFloat(personData[3], 64)
			tmp["height"] = height
		}
		if personData[4] != "" {
			is_married, _ := strconv.ParseBool(personData[4])
			tmp["isMarried"] = is_married
		}
		persons = append(persons, tmp)
	}
	return persons // TODO: replace this
}

func main() {
	data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	fmt.Println(PopulationData(data))
}
