package main

import (
	"fmt"

	"github.com/RyzhAlexWork/gointern/Task4/pkg/coffee_machine"
	"github.com/RyzhAlexWork/gointern/Task4/pkg/command"
	"github.com/RyzhAlexWork/gointern/Task4/pkg/compressor"
	"github.com/RyzhAlexWork/gointern/Task4/pkg/heater"
	"github.com/RyzhAlexWork/gointern/Task4/pkg/invoker"
)

func main() {
	newInvoker := invoker.NewInvoker()
	newCompressor := compressor.NewCompressor()
	newHeater := heater.NewHeater()
	newCoffeeMachine := coffeemachine.NewCoffeeMachine(newCompressor, newHeater)
	newCustomizationCommand := command.NewCustomizationCommand("Espresso", newCoffeeMachine)
	newIncreasePressureCommand := command.NewIncreasePressureCommand(3, newCoffeeMachine)
	newReduceTemperatureCommand := command.NewReduceTemperatureCommand(15, newCoffeeMachine)
	newInvoker.StoreCommand(newCustomizationCommand)
	newInvoker.StoreCommand(newIncreasePressureCommand)
	newInvoker.StoreCommand(newReduceTemperatureCommand)
	newInvoker.Execute()
	fmt.Println("Pressure value:", newCompressor.GiveValuePressure())
	fmt.Println("Temperature:", newHeater.GiveValueTemperature())
	newInvoker.UnStoreCommand()
}
