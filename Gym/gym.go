package Gym

type Member interface {
	GetDetails() string
}

type BasicMember struct {
	Name   string
	Plan   string
	Active bool
}

type PremiumMember struct {
	Name            string
	Plan            string
	Active          bool
	PersonalTrainer bool
}

func (b BasicMember) GetDetails() string {
	status := "inactive"
	if b.Active {
		status = "active"
	}
	return "Basic Member: " + b.Name + ", status: " + status
}
