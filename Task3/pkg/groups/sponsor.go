package groups

// Sponsor ...
type Sponsor interface {
	Accept(v visitor)
	GetLogin() string
	GetReputation() int
	AddReputation(amount int)
	ReduceReputation(amount int)
}

type sponsor struct {
	login      string
	reputation reputation
}

// Accept takes the right method from visitor.
func (s *sponsor) Accept(v visitor) {
	v.VisitSponsor(s)
}

// GetLogin get account login.
func (s *sponsor) GetLogin() string {
	return s.login
}

// GetReputation get account reputation.
func (s *sponsor) GetReputation() int {
	return s.reputation.GetAmount()
}

// AddReputation add account reputation.
func (s *sponsor) AddReputation(amount int) {
	s.reputation.Add(amount)
}

// ReduceReputation reduce account reputation.
func (s *sponsor) ReduceReputation(amount int) {
	s.reputation.Reduce(amount)
}

// NewSponsor create user implementation for interface Sponsor.
func NewSponsor(login string, inputReputation reputation) Sponsor {
	return &sponsor{
		login:      login,
		reputation: inputReputation,
	}
}
