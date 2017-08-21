package field

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math"
)

type Field struct {
    width int
	cols int
	colWidth int
	lineWidth float64
	Imd *imdraw.IMDraw
}

func Create(width int, cols int, colWidth int, color pixel.RGBA) Field {
	f := Field{}
	f.width = width
	f.cols = cols
	f.colWidth = colWidth
	f.lineWidth = math.Max(float64(colWidth) / 10, 2)
	f.Imd = imdraw.New(nil)
	f.Imd.Color = color
	for i := 0; i <= f.cols; i++ {
		f.Imd.Push(pixel.V(float64(i * colWidth), 0), pixel.V(float64(i * colWidth), float64(width)))
		f.Imd.Line(f.lineWidth)
		f.Imd.Push(pixel.V(0, float64(i * colWidth)), pixel.V(float64(width), float64(i * colWidth)))
		f.Imd.Line(f.lineWidth)
	}
	return f
}