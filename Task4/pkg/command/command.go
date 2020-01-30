package command

type coffeeMachine interface {
	Customization(coffeeName string) (done bool)
	IncreasePressure(bar int) (done bool)
	ReducePressure(bar int) (done bool)
	IncreaseTemperature(degree int) (done bool)
	ReduceTemperature(degree int) (done bool)
}

// Command ...
type Command interface {
	Execute() (done bool)
}
