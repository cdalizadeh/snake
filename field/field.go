package field

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func Create(width int, cols int, colWidth int) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)
	for i := 0; i <= cols; i++ {
		imd.Push(pixel.V(float64(i * colWidth), 0), pixel.V(float64(i * colWidth), float64(width)))
		imd.Line(5)
		imd.Push(pixel.V(0, float64(i * colWidth)), pixel.V(float64(width), float64(i * colWidth)))
		imd.Line(5)
	}
	return imd
}