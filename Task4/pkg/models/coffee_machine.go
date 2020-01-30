package models

// TemperatureLimit ...
type TemperatureLimit = int

// TemperatureLimit states
const (
	MaxTemperature TemperatureLimit = 100
	MinTemperature TemperatureLimit = 0
)

// PressureLimit ...
type PressureLimit = int

// PressureLimit states
const (
	MaxPressure TemperatureLimit = 10
	MinPressure TemperatureLimit = 0
)

// CoffeeName ...
type CoffeeName = string

// CoffeeName states
const (
	EspressoName       CoffeeName = "Espresso"
	DoubleEspressoName CoffeeName = "Double Espresso"
	LatteName          CoffeeName = "Latte"
)

// CoffeeTemperature ...
type CoffeeTemperature = int

// CoffeeTemperature states
const (
	EspressoTemp       CoffeeTemperature = 88
	DoubleEspressoTemp CoffeeTemperature = 92
	LatteTemp          CoffeeTemperature = 75
)

// CoffeePressure ...
type CoffeePressure = int

// CoffeePressure states
const (
	EspressoPres       CoffeePressure = 4
	DoubleEspressoPres CoffeePressure = 3
	LattePres          CoffeePressure = 2
)
