package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	locationDayMapping := map[string]string{
		"JKT": "senin selasa rabu kamis jumat sabtu",
		"BDG": "rabu kamis sabtu",
		"BKS": "selasa kamis jumat",
		"DPK": "senin selasa",
	}

	result := make(map[string]float32)
	for _, val := range data {
		dataDetail := strings.Split(val, ":")
		if strings.Contains(locationDayMapping[dataDetail[3]], day) {
			price, _ := strconv.ParseFloat(dataDetail[2], 64)
			if day == "senin" || day == "rabu" || day == "jumat" {
				price *= 1.1
			} else {
				price *= 1.05
			}
			fullName := fmt.Sprintf("%s-%s", dataDetail[0], dataDetail[1])
			result[fullName] = float32(price)
		}
	}
	return result // TODO: replace this
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
