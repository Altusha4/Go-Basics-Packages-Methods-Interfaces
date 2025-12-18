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
		Employee.PartTime{
			HourlyRate:  2000,
			HoursWorked: 80,
		},
		Employee.Contractor{
			ProjectRate:       50000,
			ProjectsCompleted: 4,
		},
		Employee.Intern{
			DailyRate:  8000,
			DaysWorked: 22,
		},
	}

	for _, e := range employees {
		fmt.Println("Salary:", e.CalculateSalary())
		fmt.Println("Work hours:", e.GetWorkHours())
		fmt.Println("---")
	}
}
