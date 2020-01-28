package groups

// User ...
type User interface {
	Accept(v visitor)
	GetLogin() string
	GetReputation() int
	AddReputation(amount int)
	ReduceReputation(amount int)
}

type user struct {
	login      string
	reputation reputation
}

// Accept takes the right method from visitor.
func (u *user) Accept(v visitor) {
	v.VisitUser(u)
}

// GetLogin get account login.
func (u *user) GetLogin() string {
	return u.login
}

// GetReputation get account reputation.
func (u *user) GetReputation() int {
	return u.reputation.GetAmount()
}

// AddReputation add account reputation.
func (u *user) AddReputation(amount int) {
	u.reputation.Add(amount)
}

// ReduceReputation reduce account reputation.
func (u *user) ReduceReputation(amount int) {
	u.reputation.Reduce(amount)
}

// NewUser create user implementation for interface User.
func NewUser(login string, inputReputation reputation) User {
	return &user{
		login:      login,
		reputation: inputReputation,
	}
}
