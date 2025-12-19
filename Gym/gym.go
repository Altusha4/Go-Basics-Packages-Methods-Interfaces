package Gym

import "fmt"

type Member interface {
	GetDetails() string
}

type BasicMember struct {
	Name   string
	Active bool
}

type PremiumMember struct {
	Name            string
	Active          bool
	PersonalTrainer bool
}

func (b BasicMember) GetDetails() string {
	if b.Active {
		return "Basic Member: " + b.Name + " (active)"
	}
	return "Basic Member: " + b.Name + " (inactive)"
}

func (p PremiumMember) GetDetails() string {
	trainer := ""
	if p.PersonalTrainer {
		trainer = " with trainer"
	}

	if p.Active {
		return "Premium Member: " + p.Name + " (active)" + trainer
	}
	return "Premium Member: " + p.Name + " (inactive)" + trainer
}

type Gym struct {
	Members map[uint64]Member
}

func NewGym() *Gym {
	return &Gym{Members: map[uint64]Member{}}
}

func (g *Gym) AddMember(id uint64, m Member) {
	g.Members[id] = m
}

func (g *Gym) ListMembers() {
	for id, m := range g.Members {
		fmt.Println(id, "-", m.GetDetails())
	}
}
