package main

import (
	"fmt"

	"Assignment1/Employee"
	"Assignment1/Gym"
	"Assignment1/Hotel"
	"Assignment1/Wallet"
)

func main() {
	for {
		fmt.Print(`
=== Menu ===
1. Hotel
2. Employee
3. Gym
4. Wallet
0. Exit
Choose: `)

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			demoHotel()
			wait()
		case 2:
			demoEmployee()
			wait()
		case 3:
			demoGym()
			wait()
		case 4:
			demoWallet()
			wait()
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

// ===== Hotel =====
func demoHotel() {
	fmt.Println("\n--- Hotel ---")

	h := Hotel.NewHotel()
	h.AddRoom("101", "Single", 100)
	h.AddRoom("102", "Double", 150)

	h.CheckIn("101")
	h.ListVacantRooms()
}

// ===== Employee =====
func demoEmployee() {
	fmt.Println("\n--- Employee ---")

	employees := []Employee.Employee{
		Employee.FullTime{Salary: 300000, Bonus: 0.1},
		Employee.PartTime{Rate: 2000, Hours: 80},
		Employee.Contractor{Rate: 50000, Projects: 4},
		Employee.Intern{Rate: 8000, Days: 22},
	}

	for i, e := range employees {
		fmt.Println("Employee", i+1)
		fmt.Println("Salary:", e.CalculateSalary())
		fmt.Println("Work hours:", e.GetWorkHours())
	}
}

// ===== Gym =====
func demoGym() {
	fmt.Println("\n--- Gym ---")

	g := Gym.NewGym()

	g.AddMember(1, Gym.BasicMember{
		Name:   "Alex",
		Active: true,
	})

	g.AddMember(2, Gym.PremiumMember{
		Name:            "Maria",
		Active:          true,
		PersonalTrainer: true,
	})

	g.ListMembers()
}

// ===== Wallet =====
func demoWallet() {
	fmt.Println("\n--- Wallet ---")

	w := Wallet.NewWallet()

	for {
		fmt.Print(`
Wallet Menu:
1. Add money
2. Spend money
3. Show balance
4. Show transactions
0. Back
Choose: `)

		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			return
		}

		if choice == 1 {
			var amount float64
			fmt.Print("Amount: ")
			fmt.Scanln(&amount)
			w.AddMoney(amount)
		}

		if choice == 2 {
			var amount float64
			fmt.Print("Amount: ")
			fmt.Scanln(&amount)
			w.SpendMoney(amount)
		}

		if choice == 3 {
			fmt.Println("Balance:", w.GetBalance())
		}

		if choice == 4 {
			fmt.Println("Transactions:")
			w.ShowTransactions()
		}
	}
}

func wait() {
	fmt.Print("\nPress Enter to continue...")
	fmt.Scanln()
}
