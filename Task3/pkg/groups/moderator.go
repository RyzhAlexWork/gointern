package groups

// Moderator ...
type Moderator interface {
	Accept(v visitor)
	GetLogin() string
	GetReputation() int
	AddReputation(amount int)
	ReduceReputation(amount int)
}

type moderator struct {
	login      string
	reputation reputation
}

// Accept takes the right method from visitor.
func (m *moderator) Accept(v visitor) {
	v.VisitModerator(m)
}

// GetLogin get account login.
func (m *moderator) GetLogin() string {
	return m.login
}

// GetReputation get account reputation.
func (m *moderator) GetReputation() int {
	return m.reputation.GetAmount()
}

// AddReputation add account reputation.
func (m *moderator) AddReputation(amount int) {
	m.reputation.Add(amount)
}

// ReduceReputation reduce account reputation.
func (m *moderator) ReduceReputation(amount int) {
	m.reputation.Reduce(amount)
}

// NewModerator create user implementation for interface Moderator.
func NewModerator(login string, inputReputation reputation) Moderator {
	return &moderator{
		login:      login,
		reputation: inputReputation,
	}
}
