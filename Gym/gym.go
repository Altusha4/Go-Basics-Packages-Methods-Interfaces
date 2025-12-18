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
