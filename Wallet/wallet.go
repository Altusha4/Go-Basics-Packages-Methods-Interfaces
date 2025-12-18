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
