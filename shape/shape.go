package Shapes

import (
	"fmt"
	"image"
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

func (m ShapeContainer) Render() {
	var images []*image.Paletted
	var delays []int

	for i, s := range m.Shapes {
		fmt.Println(fmt.Sprint("rendering frame: ", i, " of ", len(m.Shapes)))
		s.Update(m.Context)
		img := m.Context.Image()
		bounds := img.Bounds()
		dst := image.NewPaletted(bounds, Plan9)
		draw.Draw(dst, bounds, img, bounds.Min, draw.Src)
		images = append(images, dst)
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
