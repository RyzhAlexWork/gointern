package compressor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expect = 2
)

const (
	compressorTest = "Compressor test"
)

func Test_CoffeeMachine(t *testing.T) {
	newCompressor := NewCompressor()
	newCompressor.Increase(5)
	newCompressor.Reduce(3)

	t.Run(compressorTest, func(t *testing.T) {
		assert.Equal(t, expect, newCompressor.GiveValuePressure())
	})
}
