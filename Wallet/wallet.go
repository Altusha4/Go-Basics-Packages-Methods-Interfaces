package Wallet

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
