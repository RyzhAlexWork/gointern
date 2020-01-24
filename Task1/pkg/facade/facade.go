package facade

import "github.com/RyzhAlexWork/go-intern/Task1/pkg/models"

// User ...
type User interface {
	Add(money int) (walletStatus models.Status)
	Pay(money int) (walletStatus models.Status)
	Balance() int
}

type wallet interface {
	Pay(amount int) (done bool)
	Add(amount int) (done bool)
	Balance() int
}

type walletStatus interface {
	Get() models.Status
	Change(newText models.Status)
}

type user struct {
	login        string
	wallet       wallet
	walletStatus walletStatus
}

// Add deposit to wallet and change the walletStatus text
func (u *user) Add(money int) (walletStatus models.Status) {
	var (
		check bool
	)
	check = u.wallet.Add(money)
	if check {
		u.walletStatus.Change(models.AddSuccess)
	} else {
		u.walletStatus.Change(models.AddFail)
	}
	walletStatus = u.walletStatus.Get()
	return
}

// Add makes payment from wallet and change the walletStatus text
func (u *user) Pay(money int) (walletStatus models.Status) {
	var (
		check bool
	)
	check = u.wallet.Pay(money)
	if check {
		u.walletStatus.Change(models.PaySuccess)
	} else {
		u.walletStatus.Change(models.PayFail)
	}
	walletStatus = u.walletStatus.Get()
	return
}

// Balance show balance
func (u *user) Balance() int {
	return u.wallet.Balance()
}

// NewUser create user implementation for interface User
func NewUser(login string, inputWallet wallet, inputWalletStatus walletStatus) User {
	return &user{
		login:        login,
		wallet:       inputWallet,
		walletStatus: inputWalletStatus,
	}
}
