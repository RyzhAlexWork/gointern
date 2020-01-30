package command

type reducePressureCommand struct {
	bar           int
	coffeeMachine coffeeMachine
}

// Execute calls the command ReducePressure.
func (r *reducePressureCommand) Execute() (done bool) {
	done = r.coffeeMachine.ReducePressure(r.bar)
	return
}

// NewReducePressureCommand create reducePressureCommand implementation for interface Command.
func NewReducePressureCommand(inputBar int, inputCoffeeMachine coffeeMachine) Command {
	return &reducePressureCommand{
		bar:           inputBar,
		coffeeMachine: inputCoffeeMachine,
	}
}
