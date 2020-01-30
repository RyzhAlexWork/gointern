package heater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expect = 20
)

const (
	heaterTest = "Heater test"
)

func Test_CoffeeMachine(t *testing.T) {
	newHeater := NewHeater()
	newHeater.Increase(50)
	newHeater.Reduce(30)

	t.Run(heaterTest, func(t *testing.T) {
		assert.Equal(t, expect, newHeater.GiveValueTemperature())
	})
}
