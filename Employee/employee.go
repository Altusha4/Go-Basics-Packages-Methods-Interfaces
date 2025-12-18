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

func (p PartTime) CalculateSalary() float64 {
	return float64(p.HoursWorked) * p.HourlyRate
}

func (p PartTime) CalculateBonus() float64 {
	return 0
}

func (p PartTime) GetWorkHours() int {
	return p.HoursWorked
}

func (c Contractor) CalculateSalary() float64 {
	return float64(c.ProjectsCompleted) * c.ProjectRate
}

func (c Contractor) CalculateBonus() float64 {
	if c.ProjectsCompleted > 3 {
		return c.ProjectRate * 0.1
	}
	return 0
}

func (c Contractor) GetWorkHours() int {
	return c.ProjectsCompleted * 40
}
