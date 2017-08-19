package food

import (
	"github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel"
	//"fmt"
)

var colWidth int

type Food struct {
    Xpos int
    Ypos int
	Imd *imdraw.IMDraw
}

func (f *Food) Draw() {
	f.Imd.Reset()
	f.Imd.Clear()
	f.Imd.Color = pixel.RGB(0, 1, 1)
	f.Imd.Push(pixel.V(float64(f.Xpos * colWidth), float64(f.Ypos * colWidth)), pixel.V(float64((f.Xpos + 1) * colWidth), float64((f.Ypos + 1) * colWidth)))
	f.Imd.Rectangle(0)
}

func (f *Food) Set(x int, y int) {
	f.Xpos = x
	f.Ypos = y
	f.Draw()
}

func Init(wid int) {
    colWidth = wid
}

func Create(Xpos int, Ypos int) Food{
    f := Food{}
    f.Xpos = Xpos
    f.Ypos = Ypos
	f.Imd = imdraw.New(nil)
	f.Draw()
    return f
}