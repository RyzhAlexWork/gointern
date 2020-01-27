package main

import (
	"fmt"
	"github.com/RyzhAlexWork/gointern/Task2/pkg/car"

	"github.com/RyzhAlexWork/gointern/Task2/pkg/builder"
)

func main() {
	newCar := car.NewCar()
	carBuilder := builder.NewBuild(newCar)

	carBuilder.ConstructSportcar()
	descriptionCar := newCar.GetDescription()
	fmt.Printf("%s\n", descriptionCar)
}
