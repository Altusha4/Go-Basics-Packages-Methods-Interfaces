package Wallet

import "fmt"

type Wallet struct {
	Balance      float64
	Transactions []float64
}

func NewWallet() *Wallet {
	return &Wallet{
		Balance:      0,
		Transactions: []float64{},
	}
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}

func (w *Wallet) AddMoney(amount float64) {
	if amount <= 0 {
		return
	}
	w.Balance += amount
	w.Transactions = append(w.Transactions, amount)
}

func (w *Wallet) SpendMoney(amount float64) {
	if amount <= 0 {
		return
	}
	if amount > w.Balance {
		return
	}
	w.Balance -= amount
	w.Transactions = append(w.Transactions, -amount)
}

func (w *Wallet) ShowTransactions() {
	if len(w.Transactions) == 0 {
		fmt.Println("No transactions yet")
		return
	}

	fmt.Println("Transactions:")
	for i, t := range w.Transactions {
		fmt.Printf("%d) %.2f\n", i+1, t)
	}
}
