package command

import (
	"testing"

	"github.com/RyzhAlexWork/gointern/Task4/pkg/coffee_machine"
	"github.com/RyzhAlexWork/gointern/Task4/pkg/compressor"
	"github.com/RyzhAlexWork/gointern/Task4/pkg/heater"
	"github.com/RyzhAlexWork/gointern/Task4/pkg/invoker"
	"github.com/RyzhAlexWork/gointern/Task4/pkg/models"
	"github.com/stretchr/testify/assert"
)

var (
	expectsCustomizationCommand = []bool{true, true, true, false}
	expectsIncreaseCommand      = []bool{true, true, false, true, true, false}
	expectsReduceCommand        = []bool{true, true, false, true, true, false}
)

const (
	customizationCommandTest = "CustomizationCommand test"
	increaseCommandTest      = "IncreaseCommand Test"
	reduceCommandTest        = "ReduceCommand test"
)

func Test_Command(t *testing.T) {
	newInvoker := invoker.NewInvoker()
	newCompressor := compressor.NewCompressor()
	newHeater := heater.NewHeater()
	newCoffeeMachine := coffeemachine.NewCoffeeMachine(newCompressor, newHeater)

	// CustomizationCommand test.
	newCustomizationCommandEspresso := NewCustomizationCommand(models.EspressoName, newCoffeeMachine)
	newInvoker.StoreCommand(newCustomizationCommandEspresso)
	newCustomizationCommandDoubleEspresso := NewCustomizationCommand(models.DoubleEspressoName, newCoffeeMachine)
	newInvoker.StoreCommand(newCustomizationCommandDoubleEspresso)
	newCustomizationCommandLatteName := NewCustomizationCommand(models.LatteName, newCoffeeMachine)
	newInvoker.StoreCommand(newCustomizationCommandLatteName)
	newCustomizationCommandMocco := NewCustomizationCommand("Mocco", newCoffeeMachine)
	newInvoker.StoreCommand(newCustomizationCommandMocco)
	resultCustomizationCommand := newInvoker.Execute()
	t.Run(customizationCommandTest, func(t *testing.T) {
		for i, expect := range expectsCustomizationCommand {
			assert.Equal(t, expect, resultCustomizationCommand[i])
		}
	})

	// IncreaseCommand test.
	newInvoker = invoker.NewInvoker()
	arrayIncreaseCommand := []Command{
		NewIncreasePressureCommand(1, newCoffeeMachine),
		NewIncreasePressureCommand(2, newCoffeeMachine),
		NewIncreasePressureCommand(100, newCoffeeMachine),
		NewIncreaseTemperatureCommand(3, newCoffeeMachine),
		NewIncreaseTemperatureCommand(5, newCoffeeMachine),
		NewIncreaseTemperatureCommand(100, newCoffeeMachine),
	}
	for _, command := range arrayIncreaseCommand {
		newInvoker.StoreCommand(command)
	}
	resultIncreaseCommand := newInvoker.Execute()
	t.Run(increaseCommandTest, func(t *testing.T) {
		for i, expect := range expectsIncreaseCommand {
			assert.Equal(t, expect, resultIncreaseCommand[i])
		}
	})

	// ReduceCommand test.
	newInvoker = invoker.NewInvoker()
	arrayReduceCommand := []Command{
		NewReducePressureCommand(1, newCoffeeMachine),
		NewReducePressureCommand(2, newCoffeeMachine),
		NewReducePressureCommand(100, newCoffeeMachine),
		NewReduceTemperatureCommand(3, newCoffeeMachine),
		NewReduceTemperatureCommand(5, newCoffeeMachine),
		NewReduceTemperatureCommand(100, newCoffeeMachine),
	}
	for _, command := range arrayReduceCommand {
		newInvoker.StoreCommand(command)
	}
	resultReduceCommand := newInvoker.Execute()
	t.Run(reduceCommandTest, func(t *testing.T) {
		for i, expect := range expectsReduceCommand {
			assert.Equal(t, expect, resultReduceCommand[i])
		}
	})
}
