package builder

type car interface {
	AddWheels()
	AddCarcase()
	AddEngine()
	AddAccelerator()
	AddGun()
	GetDescription() string
}

// Builder ...
type Builder interface {
	ConstructCar()
	ConstructSportcar()
	ConstructTank()
}

type builder struct {
	car car
}

// Construct car.
func (b *builder) ConstructCar() {
	b.car.AddWheels()
	b.car.AddCarcase()
	b.car.AddEngine()
}

// Construct sport-car.
func (b *builder) ConstructSportcar() {
	b.car.AddWheels()
	b.car.AddCarcase()
	b.car.AddEngine()
	b.car.AddAccelerator()
}

// Construct tank.
func (b *builder) ConstructTank() {
	b.car.AddWheels()
	b.car.AddCarcase()
	b.car.AddEngine()
	b.car.AddGun()
}

// NewBuild create builder implementation for interface Builder
func NewBuild(inputCar car) Builder {
	return &builder{
		car: inputCar,
	}
}
