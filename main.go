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

	Noise.GenerateTerrain(1000, 1000, 1.5, 1.5, 20, rand.Int63(), 400, *dc)
}
