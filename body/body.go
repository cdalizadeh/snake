package body

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math"
	//"fmt"
)

var xpos int = 0
var ypos int = 0
var direction int = 0
var nextDir int = 0
var colWidth int
var imd *imdraw.IMDraw = imdraw.New(nil)

func Create(width int) *imdraw.IMDraw {
	colWidth = width
	redraw()
	return imd
}

func SetDir(dir int) {
	nextDir = dir
}

func changePos(dir int){
	if dir == 0 {
		xpos++
	} else if dir == 1 {
		ypos++
	} else if dir == 2 {
		xpos--
	} else if dir == 3 {
		ypos--
	}
}

func Move(){
	if direction != nextDir && math.Abs(float64(direction - nextDir)) != 2 {
		direction = nextDir
	}
	changePos(direction)
	check()
	redraw()
}

func redraw() {
	clear()
	imd.Push(pixel.V(float64(xpos * colWidth), float64(ypos * colWidth)), pixel.V(float64((xpos + 1) * colWidth), float64((ypos + 1) * colWidth)))
	imd.Rectangle(0)
}

func clear() {
	imd.Reset()
	imd.Clear()
	imd.Color = pixel.RGB(1, 1, 1)
}

func check() {
}

func kill() {
	panic("is kill")
}