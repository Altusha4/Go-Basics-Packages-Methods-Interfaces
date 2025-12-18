package Employee

type Employee interface {
	CalculateSalary() float64
	GetWorkHours() int
}

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
}

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}
