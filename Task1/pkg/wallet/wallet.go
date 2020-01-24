package wallet

// Wallet ...
type Wallet interface {
	Pay(amount int) (done bool)
	Add(amount int) (done bool)
	Balance() int
}

type wallet struct {
	money    int
	maxLimit int
	minLimit int
}

// Add deposit to wallet
func (w *wallet) Add(amount int) (done bool) {
	if w.money+amount <= w.maxLimit {
		w.money = w.money + amount
		done = true
	}
	return
}

// Pay makes payment from wallet
func (w *wallet) Pay(amount int) (done bool) {
	if w.money-amount >= w.minLimit {
		w.money = w.money - amount
		done = true
	}
	return
}

// Balance show balance
func (w *wallet) Balance() int {
	return w.money
}

// NewWallet create wallet implementation for interface Wallet
func NewWallet(minLimit int, maxLimit int) Wallet {
	return &wallet{
		minLimit: minLimit,
		maxLimit: maxLimit,
	}
}
