package models

// Status ...
type Status string

// Status states
const (
	AddSuccess Status = "Deposit was successful."
	AddFail    Status = "Deposit error. Too much money."
	PaySuccess Status = "Pay was successful."
	PayFail    Status = "Pay error. Not enough money."
)
