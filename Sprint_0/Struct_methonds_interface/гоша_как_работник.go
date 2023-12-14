package main

type Employee struct {
	name     string
	position string
	salary   int
	bonus    int
}

func (e Employee) CalculateTotalSalary() int {
	return e.salary + e.bonus
}
