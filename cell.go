package main

import (
	"github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel"
)

type Cell struct {
    Xpos int
    Ypos int
    imd *imdraw.IMDraw
}

func (c *Cell) Draw() {
	c.imd.Push(pixel.V(float64(c.Xpos) * colWidth, float64(c.Ypos) * colWidth), pixel.V(float64(c.Xpos + 1) * colWidth, float64(c.Ypos + 1) * colWidth))
	c.imd.Rectangle(0)
}

func createCell(Xpos int, Ypos int, imd *imdraw.IMDraw) Cell{
    c := Cell{}
    c.Xpos = Xpos
    c.Ypos = Ypos
    c.imd = imd
    return c
}