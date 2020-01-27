package groups

// Administrator ...
type Administrator interface {
	Accept(v visitor)
	GetLogin() string
	GetReputation() int
	AddReputation(amount int)
	ReduceReputation(amount int)
}

type administrator struct {
	login      string
	reputation reputation
}

// Accept takes the right method from visitor.
func (a *administrator) Accept(v visitor) {
	v.VisitAdministrator(a)
}

// GetLogin get account login.
func (a *administrator) GetLogin() string {
	return a.login
}

// GetReputation get account reputation.
func (a *administrator) GetReputation() int {
	return a.reputation.GetAmount()
}

// AddReputation add account reputation.
func (a *administrator) AddReputation(amount int) {
	a.reputation.Add(amount)
}

// ReduceReputation reduce account reputation.
func (a *administrator) ReduceReputation(amount int) {
	a.reputation.Reduce(amount)
}

// NewAdministrator create user implementation for interface Administrator.
func NewAdministrator(login string, inputReputation reputation) Administrator {
	return &administrator{
		login:      login,
		reputation: inputReputation,
	}
}
