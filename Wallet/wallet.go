package Wallet

import "fmt"

type Wallet struct {
	Balance      float64
	Transactions []float64
}

func NewWallet() *Wallet {
	return &Wallet{}
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	w.Transactions = append(w.Transactions, amount)
}

func (w *Wallet) SpendMoney(amount float64) {
	if amount > w.Balance {
		fmt.Println("Not enough money")
		return
	}
	w.Balance -= amount
	w.Transactions = append(w.Transactions, -amount)
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}

func (w *Wallet) ShowTransactions() {
	for _, t := range w.Transactions {
		fmt.Println(t)
	}
}
