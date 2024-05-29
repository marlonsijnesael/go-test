package Noise

import (
	Shapes "test/image/shape"

	"github.com/aquilax/go-perlin"
	"github.com/fogleman/gg"
)

type TerrainType struct {
	Name     string
	Color    Shapes.Color
	MinValue float64
}

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
var colForest = Shapes.Color{
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

var TerrainTypes = []TerrainType{
	{"snow", colSnow, 0.7},
	{"rock", colRock, 0.4},
	{"forest", colForest, 0.4},
	{"grass", colGrass, 0.3},
	{"sand", colSand, 0.2},
	{"water", colWater, 0.0001},
	{"water", colWaterDark, -100},
}

func getColor(n float64) Shapes.Color {
	for _, terrain := range TerrainTypes {
		if n > terrain.MinValue {
			return terrain.Color
		}
	}
	return TerrainTypes[len(TerrainTypes)-1].Color
}

func GenerateTerrain(width float64, height float64, alpha float64, beta float64, n int32, seed int64, res float64, dc gg.Context) {
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0.; x < width; x += 1 {
		for y := 0.; y < height; y += 1 {
			dc.DrawRectangle(x, y, 1, 1)
			col := getColor(p.Noise2D(x/res, y/res))
			dc.SetRGB255(col.R, col.G, col.B)
			dc.Fill()
		}
	}
	dc.SavePNG("out.png")

}
