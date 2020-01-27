package main

import (
	"fmt"
	"github.com/RyzhAlexWork/gointern/Task2/pkg/builder"
)

func main() {
	newCar := builder.NewCar()
	carBuilder := builder.NewBuild(newCar)

	carBuilder.ConstructSportcar()
	descriptionCar := newCar.GetDescription()
	fmt.Printf("%s\n", descriptionCar)
}
