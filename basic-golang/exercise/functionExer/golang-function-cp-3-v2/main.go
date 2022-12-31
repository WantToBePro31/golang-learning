package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	pivot, lname, shortestName := 0, 100001, ""
	for last, val := range names {
		if val == ' ' || val == ',' || val == ';' {
			tmp := ""
			for pivot < last {
				tmp += string(names[pivot])
				pivot++
			}
			if len(tmp) < lname {
				lname = len(tmp)
				shortestName = tmp
			} else if len(tmp) == lname {
				if strings.Compare(tmp, shortestName) == -1 {
					lname = len(tmp)
					shortestName = tmp
				}
			}
			pivot = last+1
		}
	}
	tmp := ""
	for pivot < len(names) {
		tmp += string(names[pivot])
		pivot++
	}
	if len(tmp) < lname {
		lname = len(tmp)
		shortestName = tmp
	}
	return shortestName // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
