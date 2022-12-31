package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	totalTicket := VIP + regular + student
	totalPrice := float32(VIP*30 + regular*20 + student*10)
	if totalPrice >= 100.0 {
		if day % 2 == 1 {
			if totalTicket < 5 {
				totalPrice *= 0.85
			} else {
				totalPrice *= 0.75
			}
		} else {
			if totalTicket < 5 {
				totalPrice *= 0.9
			} else {
				totalPrice *= 0.8
			}
		}
	}
	return totalPrice // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
