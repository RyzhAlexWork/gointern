package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expects = []string{
		"Wheels:true. Carcase:true. Engine:true. Accelerator:true. Gun:false.",
		"Wheels:true. Carcase:true. Engine:true. Accelerator:false. Gun:true.",
		"Wheels:true. Carcase:true. Engine:true. Accelerator:false. Gun:false.",
	}
	descriptionsCars []string
)

const (
	builderTest = "builderTest"
)

func Test_Builder(t *testing.T) {
	newCar := NewCar()
	build := NewBuild(newCar)
	build.ConstructSportcar()
	descriptionsCars = append(descriptionsCars, newCar.GetDescription())
	newCar = NewCar()
	build = NewBuild(newCar)
	build.ConstructTank()
	descriptionsCars = append(descriptionsCars, newCar.GetDescription())
	newCar = NewCar()
	build = NewBuild(newCar)
	build.ConstructCar()
	descriptionsCars = append(descriptionsCars, newCar.GetDescription())

	t.Run(builderTest, func(t *testing.T) {
		for i, expect := range expects {
			assert.Equal(t, expect, descriptionsCars[i])
		}
	})
}
