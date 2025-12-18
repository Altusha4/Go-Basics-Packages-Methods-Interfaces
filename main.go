package main

import (
	"Assignment1/Employee"
	"fmt"
)

func main() {
	employees := []Employee.Employee{
		Employee.FullTime{
			MonthlySalary: 300000,
			BonusRate:     0.1,
		},
	}

	for _, e := range employees {
		fmt.Println("Salary:", e.CalculateSalary())
		fmt.Println("Work hours:", e.GetWorkHours())
	}
}
