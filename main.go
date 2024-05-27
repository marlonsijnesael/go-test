package main

import (
	"math/rand"
	Shapes "test/image/shape"

	"github.com/fogleman/gg"
)

var width = 1000
var height = 1000

func main() {
	dc := gg.NewContext(width, height)
	shapeContainer := Shapes.ShapeContainer{
		Context: *dc,
	}
	for i := 0; i < 90; i++ {
		circle := Shapes.Circle{
			X: float64(i) * (50 * rand.Float64()),
			Y: float64(i) * (50 * rand.Float64()),
			R: 20 * rand.Float64(),
		}
		shapeContainer.Shapes[i] = circle
	}

	shapeContainer.Render()
}
