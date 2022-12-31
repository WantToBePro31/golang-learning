package main

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name string
	BaseSalary int
	WorkingMonth int
}

func (j Junior) GetBonus() float64 {
	var prorata float64
	if j.WorkingMonth >= 12 {
		prorata = 1.0
	} else {
		prorata = float64(j.WorkingMonth)/12.0
	}
	return float64(j.BaseSalary) * prorata
}

type Senior struct {
	Name string
	BaseSalary int
	WorkingMonth int
	PerformanceRate float64
}

func (s Senior) GetBonus() float64 {
	var prorata float64
	if s.WorkingMonth >= 12 {
		prorata = 1.0
	} else {
		prorata = float64(s.WorkingMonth)/12.0
	}
	return 2.0 * float64(s.BaseSalary) * prorata + (s.PerformanceRate * float64(s.BaseSalary))
}

type Manager struct {
	Name string
	BaseSalary int
	WorkingMonth int
	PerformanceRate float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	var prorata float64
	if m.WorkingMonth >= 12 {
		prorata = 1.0
	} else {
		prorata = float64(m.WorkingMonth)/12.0
	}
	return 2.0 * float64(m.BaseSalary) * prorata + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	var totalBonus float64
	for i := 0; i < len(employees); i++ {
		totalBonus += employees[i].GetBonus()
	}
	return totalBonus // TODO: replace this
}
