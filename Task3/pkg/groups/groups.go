package groups

type visitor interface {
	VisitAdministrator(a Administrator)
	VisitModerator(m Moderator)
	VisitSponsor(s Sponsor)
	VisitUser(u User)
}

type reputation interface {
	GetAmount() int
	Add(amount int)
	Reduce(amount int)
}

// Groups ...
type Groups interface {
	Accept(v visitor)
	GetLogin() string
	GetReputation() int
	AddReputation(amount int)
	ReduceReputation(amount int)
}
