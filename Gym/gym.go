package Gym

type Member interface {
	GetDetails() string
}

type BasicMember struct {
	Name   string
	Plan   string
	Active bool
}
