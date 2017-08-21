package body

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/cdalizadeh/snake/cell"
	"math"
)

var colWidth float64
var cols int

type Body struct {
    s []cell.Cell
	head cell.Cell
	direction int
	nextDir int
	Imd *imdraw.IMDraw
	addCell bool
	color pixel.RGBA
}

func (b *Body) Draw() {
	b.Imd.Reset()
	b.Imd.Clear()
	b.Imd.Color = b.color
	for i := 0; i < len(b.s); i++ {
		b.s[i].Draw()
	}
}

func (b *Body) SetDir(dir int) {
	if b.direction != dir && math.Abs(float64(b.direction - dir)) != 2 {
		b.nextDir = dir
	}
}

func (b *Body) ChangePos(){
	if b.addCell {
		b.addCell = false
	} else {
		b.s = b.s[1:]
	}
	if b.direction == 0 {
		b.head = cell.Create(b.head.Xpos + 1, b.head.Ypos)
	} else if b.direction == 1 {
		b.head = cell.Create(b.head.Xpos, b.head.Ypos + 1)
	} else if b.direction == 2 {
		b.head = cell.Create(b.head.Xpos - 1, b.head.Ypos)
	} else if b.direction == 3 {
		b.head = cell.Create(b.head.Xpos, b.head.Ypos - 1)
	}
	b.s = append(b.s, b.head)
}

func (b *Body) Move() {
	b.direction = b.nextDir
	b.ChangePos()
	b.Draw()
	b.Check()
}

func (b *Body) GetHead() (int, int) {
	return b.head.Xpos, b.head.Ypos
}

func (b *Body) Check() {
	for i:= 0; i < len(b.s) - 1; i++ {
		if b.head.Xpos == b.s[i].Xpos && b.head.Ypos == b.s[i].Ypos {
			kill()
		}
	}
	if b.head.Xpos < 0 {
		kill()
	}
	if b.head.Ypos < 0 {
		kill()
	}
	if b.head.Xpos > cols - 1 {
		kill()
	}
	if b.head.Ypos > cols - 1 {
		kill()
	}
}

func (b *Body) Eat() {
	b.addCell = true
}

func (b *Body) IsWithinBody(x int, y int) bool{
	for i := 0; i < len(b.s); i++ {
		if x == b.s[i].Xpos && y == b.s[i].Ypos {
			return true
		}
	}
	return false
}

func kill(){
	panic(6)
}

func Init(numCols int, width float64) {
	colWidth = width
	cols = numCols
}

func Create(x int, y int, dir int, color pixel.RGBA) Body {
	b := Body{}
	b.s = make([]cell.Cell, 1, cols * cols)
	b.Imd = imdraw.New(nil)
	b.direction = 0
	b.nextDir = 0
	b.addCell = false
	b.color = color
	cell.Init(colWidth, b.Imd)
	b.s[0] = cell.Create(0, 0)
	b.head = b.s[0]
	b.Draw()
	return b
}