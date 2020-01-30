package command

type increasePressureCommand struct {
	bar           int
	coffeeMachine coffeeMachine
}

// Execute calls the command IncreasePressure.
func (c *increasePressureCommand) Execute() (done bool) {
	done = c.coffeeMachine.IncreasePressure(c.bar)
	return
}

// NewIncreasePressureCommand create increasePressureCommand implementation for interface Command.
func NewIncreasePressureCommand(inputBar int, inputCoffeeMachine coffeeMachine) Command {
	return &increasePressureCommand{
		bar:           inputBar,
		coffeeMachine: inputCoffeeMachine,
	}
}
