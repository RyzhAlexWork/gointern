package builder

import "strconv"

// Car ...
type Car interface {
	AddWheels()
	AddCarcase()
	AddEngine()
	AddAccelerator()
	AddGun()
	GetDescription() string
}

type car struct {
	wheels      bool
	carcase     bool
	engine      bool
	accelerator bool
	gun         bool
}

// AddGun builds a wheels.
func (c *car) AddWheels() {
	c.wheels = true
}

// AddGun builds a carcase.
func (c *car) AddCarcase() {
	c.carcase = true
}

// AddGun builds a engine.
func (c *car) AddEngine() {
	c.engine = true
}

// AddGun builds a accelerator.
func (c *car) AddAccelerator() {
	c.accelerator = true
}

// AddGun builds a gun.
func (c *car) AddGun() {
	c.gun = true
}

// GetDescription get description car.
func (c *car) GetDescription() string {
	return "Wheels:" + strconv.FormatBool(c.wheels) + ". Carcase:" + strconv.FormatBool(c.carcase) +
		". Engine:" + strconv.FormatBool(c.engine) + ". Accelerator:" + strconv.FormatBool(c.accelerator) +
		". Gun:" + strconv.FormatBool(c.gun) + "."
}

// NewCar create car implementation for interface Car
func NewCar() Car {
	return &car{}
}
