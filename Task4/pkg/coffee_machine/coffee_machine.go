package coffeemachine

import "github.com/RyzhAlexWork/gointern/Task4/pkg/models"

type compressor interface {
	Increase(bar int)
	Reduce(bar int)
	GiveValuePressure() (valuePressure int)
}

type heater interface {
	Increase(degree int)
	Reduce(degree int)
	GiveValueTemperature() (valueTemperature int)
}

// CoffeeMachine ...
type CoffeeMachine interface {
	Customization(coffeeName string) (done bool)
	IncreasePressure(bar int) (done bool)
	ReducePressure(bar int) (done bool)
	IncreaseTemperature(degree int) (done bool)
	ReduceTemperature(degree int) (done bool)
}

type coffeeMachine struct {
	maxTemperature models.TemperatureLimit
	minTemperature models.TemperatureLimit
	maxPressure    models.PressureLimit
	minPressure    models.PressureLimit
	compressor     compressor
	heater         heater
}

// Customization sets up a coffee machine.
func (c *coffeeMachine) Customization(coffeeName string) (done bool) {
	temperature := c.heater.GiveValueTemperature()
	pressure := c.compressor.GiveValuePressure()

	switch coffeeName {
	case models.EspressoName:
		if temperature > models.EspressoTemp {
			c.heater.Reduce(temperature - models.EspressoTemp)
		} else {
			c.heater.Increase(models.EspressoTemp - temperature)
		}
		if pressure > models.EspressoPres {
			c.compressor.Reduce(pressure - models.EspressoPres)
		} else {
			c.compressor.Increase(models.EspressoPres - pressure)
		}
		done = true
	case models.DoubleEspressoName:
		if temperature > models.DoubleEspressoTemp {
			c.heater.Reduce(temperature - models.DoubleEspressoTemp)
		} else {
			c.heater.Increase(models.DoubleEspressoTemp - temperature)
		}
		if pressure > models.DoubleEspressoPres {
			c.compressor.Reduce(pressure - models.DoubleEspressoPres)
		} else {
			c.compressor.Increase(models.DoubleEspressoPres - pressure)
		}
		done = true
	case models.LatteName:
		if temperature > models.LatteTemp {
			c.heater.Reduce(temperature - models.LatteTemp)
		} else {
			c.heater.Increase(models.LatteTemp - temperature)
		}
		if pressure > models.LattePres {
			c.compressor.Reduce(pressure - models.LattePres)
		} else {
			c.compressor.Increase(models.LattePres - pressure)
		}
		done = true
	}
	return
}

// IncreasePressure increases pressure.
func (c *coffeeMachine) IncreasePressure(bar int) (done bool) {
	if c.compressor.GiveValuePressure()+bar <= c.maxPressure {
		c.compressor.Increase(bar)
		done = true
	}
	return
}

// ReducePressure reduce pressure.
func (c *coffeeMachine) ReducePressure(bar int) (done bool) {
	if c.compressor.GiveValuePressure()-bar >= c.minPressure {
		c.compressor.Reduce(bar)
		done = true
	}
	return
}

// IncreaseTemperature increases temperature.
func (c *coffeeMachine) IncreaseTemperature(degree int) (done bool) {
	if c.heater.GiveValueTemperature()+degree <= c.maxTemperature {
		c.heater.Increase(degree)
		done = true
	}
	return
}

// ReduceTemperature reduce temperature.
func (c *coffeeMachine) ReduceTemperature(degree int) (done bool) {
	if c.heater.GiveValueTemperature()-degree >= c.minTemperature {
		c.heater.Reduce(degree)
		done = true
	}
	return
}

// NewCoffeeMachine create coffeeMachine implementation for interface CoffeeMachine.
func NewCoffeeMachine(inputCompressor compressor, inputHeater heater) CoffeeMachine {
	return &coffeeMachine{
		maxTemperature: models.MaxTemperature,
		minTemperature: models.MinTemperature,
		maxPressure:    models.MaxPressure,
		minPressure:    models.MinPressure,
		compressor:     inputCompressor,
		heater:         inputHeater,
	}
}
