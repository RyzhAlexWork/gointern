package coffeemachine

import (
	"testing"

	compressorPackage "github.com/RyzhAlexWork/gointern/Task4/pkg/compressor"
	heaterPackage "github.com/RyzhAlexWork/gointern/Task4/pkg/heater"
	"github.com/stretchr/testify/assert"
)

var (
	expects = []bool{true, true, false, true, true, false, true, true, false, true, true, false}
)

const (
	coffeeMachineTest = "CoffeeMachine test"
)

func Test_CoffeeMachine(t *testing.T) {
	newCompressor := compressorPackage.NewCompressor()
	newHeater := heaterPackage.NewHeater()
	newCoffeeMachine := NewCoffeeMachine(newCompressor, newHeater)
	result := []bool{
		newCoffeeMachine.IncreaseTemperature(5),
		newCoffeeMachine.IncreaseTemperature(50),
		newCoffeeMachine.IncreaseTemperature(500),
		newCoffeeMachine.ReduceTemperature(3),
		newCoffeeMachine.ReduceTemperature(30),
		newCoffeeMachine.ReduceTemperature(300),
		newCoffeeMachine.IncreasePressure(1),
		newCoffeeMachine.IncreasePressure(5),
		newCoffeeMachine.IncreasePressure(20),
		newCoffeeMachine.ReducePressure(2),
		newCoffeeMachine.ReducePressure(3),
		newCoffeeMachine.ReducePressure(25),
	}

	t.Run(coffeeMachineTest, func(t *testing.T) {
		for i, expect := range expects {
			assert.Equal(t, expect, result[i])
		}
	})
}
