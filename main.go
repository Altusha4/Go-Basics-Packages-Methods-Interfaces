package main

import (
	"Assignment1/Employee"
	"Assignment1/Gym"
	"Assignment1/Hotel"
	"Assignment1/Wallet"
	"fmt"
)

func main() {
	for {
		fmt.Print(`
=== Menu ===
1. Task 1: Hotel
2. Task 2: Employee
3. Task 3: Gym
4. Task 4: Wallet
0. Exit
Choose: `)

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			demoHotel()
		case 2:
			demoEmployee()
		case 3:
			demoGym()
		case 4:
			demoWallet()
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

// Task 1: Hotel
func demoHotel() {
	fmt.Println("\n--- Task 1: Hotel ---")

	hotel := Hotel.NewHotel()
	hotel.AddRoom("101", "Single", 100)
	hotel.AddRoom("102", "Double", 150)

	hotel.CheckIn("101")
	hotel.ListVacantRooms()
}

// Task 2: Employee
func demoEmployee() {
	fmt.Println("\n--- Task 2: Employee ---")

	employees := []Employee.Employee{
		Employee.FullTime{MonthlySalary: 300000, BonusRate: 0.1},
		Employee.PartTime{HourlyRate: 2000, HoursWorked: 80},
		Employee.Contractor{ProjectRate: 50000, ProjectsCompleted: 4},
		Employee.Intern{DailyRate: 8000, DaysWorked: 22},
	}

	for _, e := range employees {
		fmt.Println("Salary:", e.CalculateSalary())
		fmt.Println("Work hours:", e.GetWorkHours())
		fmt.Println("---")
	}
}

// Task 3: Gym
func demoGym() {
	fmt.Println("\n--- Task 3: Gym ---")

	gym := Gym.NewGym()

	gym.AddMember(1, Gym.BasicMember{
		Name:   "Alex",
		Plan:   "Basic",
		Active: true,
	})

	gym.AddMember(2, Gym.PremiumMember{
		Name:            "Maria",
		Plan:            "Premium",
		Active:          true,
		PersonalTrainer: true,
	})

	gym.ListMembers()
}

// Task 4: Wallet
func demoWallet() {
	fmt.Println("\n--- Task 4: Wallet ---")

	w := Wallet.NewWallet()

	for {
		fmt.Print(`
Wallet Menu:
1. Add money
2. Spend money
3. Show balance
0. Back
Choose: `)

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Amount: ")
			fmt.Scanln(&amount)
			w.AddMoney(amount)

		case 2:
			var amount float64
			fmt.Print("Amount: ")
			fmt.Scanln(&amount)
			w.SpendMoney(amount)

		case 3:
			fmt.Println("Balance:", w.GetBalance())

		case 0:
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
