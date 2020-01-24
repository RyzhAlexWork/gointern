package builder

// Builder ...
type Builder interface {
	ConstructCar()
	ConstructSportcar()
	ConstructTank()
}

type builder struct {
	car Car
}

// Construct car.
func (b *builder) ConstructCar() {
	b.car.AddWheels()
	b.car.AddCarcase()
	b.car.AddEngine()
	b.car.MakeDescription()
}

// Construct sport-car.
func (b *builder) ConstructSportcar() {
	b.car.AddWheels()
	b.car.AddCarcase()
	b.car.AddEngine()
	b.car.AddAccelerator()
	b.car.MakeDescription()
}

// Construct tank.
func (b *builder) ConstructTank() {
	b.car.AddWheels()
	b.car.AddCarcase()
	b.car.AddEngine()
	b.car.AddGun()
	b.car.MakeDescription()
}

// NewBuild create builder implementation for interface Builder
func NewBuild(inputCar Car) Builder {
	return &builder{
		car: inputCar,
	}
}
