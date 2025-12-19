package Employee

type Employee interface {
	CalculateSalary() float64
	GetWorkHours() int
}

type FullTime struct {
	Salary float64
	Bonus  float64
}

type PartTime struct {
	Rate  float64
	Hours int
}

type Contractor struct {
	Rate     float64
	Projects int
}

type Intern struct {
	Rate float64
	Days int
}

func (f FullTime) CalculateSalary() float64 {
	return f.Salary + f.Salary*f.Bonus
}

func (f FullTime) GetWorkHours() int {
	return 160
}

func (p PartTime) CalculateSalary() float64 {
	return float64(p.Hours) * p.Rate
}

func (p PartTime) GetWorkHours() int {
	return p.Hours
}

func (c Contractor) CalculateSalary() float64 {
	return float64(c.Projects) * c.Rate
}

func (c Contractor) GetWorkHours() int {
	return c.Projects * 40
}

func (i Intern) CalculateSalary() float64 {
	return float64(i.Days) * i.Rate
}

func (i Intern) GetWorkHours() int {
	return i.Days * 8
}
