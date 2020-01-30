package heater

// Heater ...
type Heater interface {
	Increase(degree int)
	Reduce(degree int)
	GiveValueTemperature() (valueTemperature int)
}

type heater struct {
	degree int
}

// Increase value degree.
func (h *heater) Increase(degree int) {
	h.degree += degree
}

// Reduce value degree.
func (h *heater) Reduce(degree int) {
	h.degree -= degree
}

// GiveValueTemperature return value temperature.
func (h *heater) GiveValueTemperature() (valueTemperature int) {
	valueTemperature = h.degree
	return
}

// NewHeater create heater implementation for interface Heater.
func NewHeater() Heater {
	return &heater{}
}
