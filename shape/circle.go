package Shapes

import (
	"github.com/fogleman/gg"
)

type Circle struct {
	R   float64
	X   float64
	Y   float64
	Col Color
}

func (c Circle) Update(dc gg.Context) {
	dc.DrawCircle(c.X, c.Y, c.R)
	dc.SetRGBA(c.Col.R, c.Col.G, c.Col.B, c.Col.B)
	dc.Fill()
}
