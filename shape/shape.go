package Shapes

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"

	"github.com/fogleman/gg"
)

type Shape interface {
	Update(dc gg.Context)
}

type ShapeContainer struct {
	Context gg.Context
	Shapes  [90]Shape
}

var palette color.Palette = color.Palette{
	image.Transparent,
	color.RGBA{0xFF, 0x00, 0x00, 255},
}

func (m ShapeContainer) Render() {
	var images []*image.Paletted
	var delays []int

	for i, s := range m.Shapes {
		fmt.Println(i)
		s.Update(m.Context)
		img := m.Context.Image()
		bounds := img.Bounds()

		dst := image.NewPaletted(bounds, palette)
		draw.Draw(dst, bounds, img, bounds.Min, draw.Src)
		images = append(images, dst)
		fmt.Println(images)
		delays = append(delays, 10)

	}

	f, err := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
		//Disposal: disposals,
	})

}
