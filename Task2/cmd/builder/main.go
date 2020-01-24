package main

import (
	"fmt"
	"github.com/RyzhAlexWork/gointern/Task2/pkg/builder"
)

func main() {
	newCar := builder.NewCar()
	build := builder.NewBuild(newCar)

	build.ConstructSportcar()
	descriptionCar := newCar.GetDescription()
	fmt.Printf("%s\n", descriptionCar)
}
