package main

import (
	"math/rand"
	Noise "test/image/noise"

	"github.com/fogleman/gg"
)

var width = 1000
var height = 1000

func main() {
	dc := gg.NewContext(width, height)

	Noise.GetNoiseMap(1.5, 1.5, 20, rand.Int63(), 400, *dc)

	// shapeContainer := Shapes.ShapeContainer{
	// 	Context: *dc,
	// }
	// for i := 0; i < len(shapeContainer.Shapes); i++ {
	// 	circle := Shapes.Circle{
	// 		X: float64(i) * (50 * rand.Float64()),
	// 		Y: float64(i) * (50 * rand.Float64()),
	// 		R: 20 * rand.Float64(),
	// 		Col: Shapes.Color{
	// 			R: rand.Float64(),
	// 			G: rand.Float64(),
	// 			B: rand.Float64(),
	// 			A: rand.Float64(),
	// 		},
	// 	}
	// 	shapeContainer.Shapes[i] = circle
	// }

	// shapeContainer.Render()
}
