package main

import "fmt"

func Add(a, b int) int {
	return a+b
}

type School struct {
	Name string
	Address string
	Grades []int
}

func (s *School) MaxMin() (int, int) {
	max, min := -1, 101
	for i := 0; i < len(s.Grades); i++ {
		if s.Grades[i] > max {
			max = s.Grades[i]
		}
		if s.Grades[i] < min {
			min = s.Grades[i]
		}
	}
	return min, max
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}
	s.AddGrade(s.Grades...)
	sum := 0
	for i := 0; i < len(s.Grades); i++ {
		sum = Add(sum, s.Grades[i])
	}
	Min, Max := s.MaxMin()
	return float64(sum)/float64(len(s.Grades)), Min, Max  // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 100, 90, 100, 90},
	})

	fmt.Println(avg, min, max)
}
