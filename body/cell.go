package body

import (
	"github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel"
)

var imd *imdraw.IMDraw

type Cell struct {
    Xpos int
    Ypos int
}

func (c *Cell) Draw() {
	imd.Push(pixel.V(float64(c.Xpos) * colWidth, float64(c.Ypos) * colWidth), pixel.V(float64(c.Xpos + 1) * colWidth, float64(c.Ypos + 1) * colWidth))
	imd.Rectangle(0)
}

func initCell(wid float64, bodyImd *imdraw.IMDraw) {
    imd = bodyImd
}

func createCell(Xpos int, Ypos int) Cell{
    c := Cell{}
    c.Xpos = Xpos
    c.Ypos = Ypos
    return c
}