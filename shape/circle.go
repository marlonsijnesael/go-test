package Shapes

import (
	"fmt"

	"github.com/fogleman/gg"
)

type Circle struct {
	R float64
	X float64
	Y float64
	// C color.RGBA
}

func (c Circle) Update(dc gg.Context) {
	fmt.Println("context:")
	fmt.Println(dc)
	dc.DrawCircle(c.X, c.Y, c.R)
	dc.SetRGB(1, 0, 0)
	dc.Fill()
}
