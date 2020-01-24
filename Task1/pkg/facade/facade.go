package facade

import "github.com/RyzhAlexWork/go-intern/Task1/pkg/models"

type wallet interface {
	Pay(amount int) (done bool)
	Add(amount int) (done bool)
	Balance() int
}

type walletStatus interface {
	Get() models.Status
	Change(newText models.Status)
}

// User ...
type User interface {
	Add(money int) (walletStatus models.Status)
	Pay(money int) (walletStatus models.Status)
	Balance() int
}

type user struct {
	login        string
	wallet       wallet
	walletStatus walletStatus
	addStatus    map[bool]models.Status
	payStatus    map[bool]models.Status
}

// Add deposit to wallet and change the walletStatus text
func (u *user) Add(money int) (walletStatus models.Status) {
	u.walletStatus.Change(u.addStatus[u.wallet.Add(money)])
	walletStatus = u.walletStatus.Get()
	return
}

// Add makes payment from wallet and change the walletStatus text
func (u *user) Pay(money int) (walletStatus models.Status) {
	u.walletStatus.Change(u.payStatus[u.wallet.Pay(money)])
	walletStatus = u.walletStatus.Get()
	return
}

// Balance show balance
func (u *user) Balance() int {
	return u.wallet.Balance()
}

// NewUser create user implementation for interface User
func NewUser(login string, inputWallet wallet, inputWalletStatus walletStatus) User {
	addStatus := make(map[bool]models.Status)
	addStatus[true] = models.AddSuccess
	addStatus[false] = models.AddFail
	payStatus := make(map[bool]models.Status)
	payStatus[true] = models.PaySuccess
	payStatus[false] = models.PayFail
	return &user{
		login:        login,
		wallet:       inputWallet,
		walletStatus: inputWalletStatus,
		addStatus:    addStatus,
		payStatus:    payStatus,
	}
}
