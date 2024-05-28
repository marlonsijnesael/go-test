package Noise

import (
	Shapes "test/image/shape"

	"github.com/aquilax/go-perlin"
	"github.com/fogleman/gg"
)

var colSnow = Shapes.Color{
	R: 235,
	G: 235,
	B: 235,
	A: 1,
}

var colWater = Shapes.Color{
	R: 3,
	G: 121,
	B: 113,
	A: 1,
}
var colWaterDark = Shapes.Color{
	R: 2,
	G: 52,
	B: 54,
	A: 1,
}

var colGrass = Shapes.Color{
	R: 24,
	G: 181,
	B: 93,
	A: 1,
}
var colGrassDark = Shapes.Color{
	R: 0,
	G: 121,
	B: 3,
	A: 1,
}

var colSand = Shapes.Color{
	R: 214,
	G: 171,
	B: 69,
	A: 1,
}

var colRock = Shapes.Color{
	R: 45,
	G: 58,
	B: 49,
	A: 1,
}

func getColor(n float64) Shapes.Color {
	if n > 0.7 {
		return colSnow
	} else if n > 0.5 {
		return colRock
	} else if n > 0.4 {
		return colGrassDark
	} else if n > 0.3 {
		return colGrass
	} else if n > 0.2 {
		return colSand
	} else if n > 0.0001 {
		return colWater
	}
	return colWaterDark
}

func GetNoiseMap(alpha float64, beta float64, n int32, seed int64, res float64, dc gg.Context) {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < 1000; x += 1 {
		for y := 0.; y < 1000; y += 1 {
			// fmt.Printf("%0.0f;%0.0f;%0.4f\n", x, y, p.Noise2D(x/10, y/10))
			dc.DrawRectangle(x, y, 1, 1)
			col := getColor(p.Noise2D(x/res, y/res))
			dc.SetRGB(col.R/255, col.G/255, col.B/255)
			dc.Fill()
		}
	}
	dc.SavePNG("out.png")

}
