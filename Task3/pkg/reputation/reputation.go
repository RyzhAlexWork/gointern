package reputation

// Reputation ...
type Reputation interface {
	GetAmount() int
	Add(amount int)
	Reduce(amount int)
}

type reputation struct {
	amount int
}

// GetAmount get amount reputation.
func (r *reputation) GetAmount() int {
	return r.amount
}

// Add add reputation.
func (r *reputation) Add(amount int) {
	r.amount += amount
}

// Reduce reduce reputation.
func (r *reputation) Reduce(amount int) {
	r.amount -= amount
}

// NewReputation create user implementation for interface Reputation.
func NewReputation(amount int) Reputation {
	return &reputation{
		amount: amount,
	}
}
