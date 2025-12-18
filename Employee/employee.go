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

type Contractor struct {
	ProjectRate       float64
	ProjectsCompleted int
}

type Intern struct {
	DailyRate  float64
	DaysWorked int
}

func (f FullTime) CalculateSalary() float64 {
	return f.MonthlySalary + f.CalculateBonus()
}

func (f FullTime) CalculateBonus() float64 {
	return f.MonthlySalary * f.BonusRate
}

func (f FullTime) GetWorkHours() int {
	return 160
}
