package command

type customizationCommand struct {
	coffeeName    string
	coffeeMachine coffeeMachine
}

// Execute calls the command Customization.
func (c *customizationCommand) Execute() (done bool) {
	done = c.coffeeMachine.Customization(c.coffeeName)
	return
}

// NewCustomizationCommand create customizationCommand implementation for interface Command.
func NewCustomizationCommand(inputCoffeeName string, inputCoffeeMachine coffeeMachine) Command {
	return &customizationCommand{
		coffeeName:    inputCoffeeName,
		coffeeMachine: inputCoffeeMachine,
	}
}
