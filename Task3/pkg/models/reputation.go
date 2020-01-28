package models

// Reputation ...
type Reputation = int

// Reputation states
const (
	ForAdministrator Reputation = 1000
	ForModerator     Reputation = 500
	ForSponsor       Reputation = 100
	ForUser          Reputation = 50
)
