package command

type increaseTemperatureCommand struct {
	degree        int
	coffeeMachine coffeeMachine
}

// Execute calls the command IncreaseTemperature.
func (c *increaseTemperatureCommand) Execute() (done bool) {
	done = c.coffeeMachine.IncreaseTemperature(c.degree)
	return
}

// NewIncreaseTemperatureCommand create increaseTemperatureCommand implementation for interface Command.
func NewIncreaseTemperatureCommand(inputDegree int, inputCoffeeMachine coffeeMachine) Command {
	return &increaseTemperatureCommand{
		degree:        inputDegree,
		coffeeMachine: inputCoffeeMachine,
	}
}
