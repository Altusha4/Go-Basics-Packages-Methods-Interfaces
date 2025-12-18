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

func (p PremiumMember) GetDetails() string {
	status := "inactive"
	if p.Active {
		status = "active"
	}

	trainer := "no personal trainer"
	if p.PersonalTrainer {
		trainer = "with personal trainer"
	}

	return "Premium Member: " + p.Name + ", status: " + status + ", " + trainer
}

type Gym struct {
	Members map[uint64]Member
}

func NewGym() *Gym {
	return &Gym{
		Members: make(map[uint64]Member),
	}
}

func (g *Gym) AddMember(id uint64, member Member) {
	g.Members[id] = member
}

func (g *Gym) ListMembers() {
	for id, member := range g.Members {
		println("ID:", id, "-", member.GetDetails())
	}
}
