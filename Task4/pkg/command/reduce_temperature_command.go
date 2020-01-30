package command

type reduceTemperatureCommand struct {
	degree        int
	coffeeMachine coffeeMachine
}

// Execute calls the command ReduceTemperature.
func (r *reduceTemperatureCommand) Execute() (done bool) {
	done = r.coffeeMachine.ReduceTemperature(r.degree)
	return
}

// NewReduceTemperatureCommand create reduceTemperatureCommand implementation for interface Command.
func NewReduceTemperatureCommand(inputDegree int, inputCoffeeMachine coffeeMachine) Command {
	return &reduceTemperatureCommand{
		degree:        inputDegree,
		coffeeMachine: inputCoffeeMachine,
	}
}
