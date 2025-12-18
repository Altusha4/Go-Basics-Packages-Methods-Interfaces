package main

import (
	"Assignment1/Gym"
)

func main() {
	gym := Gym.NewGym()

	basic := Gym.BasicMember{
		Name:   "Alex",
		Plan:   "Basic",
		Active: true,
	}

	premium := Gym.PremiumMember{
		Name:            "Maria",
		Plan:            "Premium",
		Active:          true,
		PersonalTrainer: true,
	}

	gym.AddMember(1, basic)
	gym.AddMember(2, premium)

	gym.ListMembers()
}
