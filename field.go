package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math"
)

type Field struct {
    width float64
	cols int
	colWidth float64
	lineWidth float64
	Imd *imdraw.IMDraw
	color pixel.RGBA
}

func (f *Field) Draw() {
	f.Imd.Reset()
	f.Imd.Clear()
	f.Imd.Color = f.color
	for i := 0.0; i <= float64(f.cols); i++ {
		f.Imd.Push(pixel.V(i * f.colWidth, 0), pixel.V(i * f.colWidth, f.width))
		f.Imd.Line(f.lineWidth)
		f.Imd.Push(pixel.V(0, i * f.colWidth), pixel.V(f.width, i * f.colWidth))
		f.Imd.Line(f.lineWidth)
	}
}

func createField(width float64, cols int, colWidth float64, color pixel.RGBA) Field {
	f := Field{}
	f.width = width
	f.cols = cols
	f.colWidth = colWidth
	f.lineWidth = math.Max(colWidth / 10, 2)
	f.Imd = imdraw.New(nil)
	f.color = color
	f.Draw()
	return f
}