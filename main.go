package main

import (
	"Assignment1/Wallet"
	"fmt"
)

func main() {
	w := Wallet.NewWallet()

	for {
		fmt.Print(`
=== Wallet Menu ===
1. Add money
2. Spend money
3. Show balance
0. Exit
Choose: `)

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Amount to add: ")
			fmt.Scanln(&amount)
			w.AddMoney(amount)

		case 2:
			var amount float64
			fmt.Print("Amount to spend: ")
			fmt.Scanln(&amount)
			w.SpendMoney(amount)

		case 3:
			fmt.Println("Current balance:", w.GetBalance())

		case 0:
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
