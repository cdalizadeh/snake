package main

import (
	"github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel"
)

type Food struct {
    Xpos int
    Ypos int
	Imd *imdraw.IMDraw
	color pixel.RGBA
}

func (f *Food) Draw() {
	f.Imd.Reset()
	f.Imd.Clear()
	f.Imd.Color = f.color
	f.Imd.Push(pixel.V(float64(float64(f.Xpos) * colWidth), float64(float64(f.Ypos) * colWidth)), pixel.V(float64((float64(f.Xpos) + 1) * colWidth), float64((float64(f.Ypos) + 1) * colWidth)))
	f.Imd.Rectangle(0)
}

func (f *Food) Set(x int, y int) {
	f.Xpos = x
	f.Ypos = y
	f.Draw()
}

func createFood(Xpos int, Ypos int, color pixel.RGBA) Food{
    f := Food{}
    f.Xpos = Xpos
    f.Ypos = Ypos
	f.Imd = imdraw.New(nil)
	f.color = color
	f.Draw()
    return f
}