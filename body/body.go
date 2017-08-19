package body

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/cdalizadeh/snake/cell"
	"math"
	//"fmt"
)

var colWidth int
var cols int

type Body struct {
    s []cell.Cell
	direction int
	nextDir int
	Imd *imdraw.IMDraw
	addCell bool
}

func (b *Body) Redraw() {
	b.Imd.Reset()
	b.Imd.Clear()
	b.Imd.Color = pixel.RGB(1, 1, 1)
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

	if b.direction == 0 {
		b.s = b.s[1:]
		b.s = append(b.s, cell.Create(b.s[len(b.s) - 1].Xpos + 1, b.s[len(b.s) - 1].Ypos))
	} else if b.direction == 1 {
		b.s = b.s[1:]
		b.s = append(b.s, cell.Create(b.s[len(b.s) - 1].Xpos, b.s[len(b.s) - 1].Ypos + 1))
	} else if b.direction == 2 {
		b.s = b.s[1:]
		b.s = append(b.s, cell.Create(b.s[len(b.s) - 1].Xpos - 1, b.s[len(b.s) - 1].Ypos))
	} else if b.direction == 3 {
		b.s = b.s[1:]
		b.s = append(b.s, cell.Create(b.s[len(b.s) - 1].Xpos, b.s[len(b.s) - 1].Ypos - 1))
	}
}

func (b *Body) Move(){
	b.direction = b.nextDir
	b.ChangePos()
	b.Check()
	b.Redraw()
}

func (b *Body) Check() {
	/*if b.s[0].Xpos < 0 {
		b.s[0].Xpos = 0
	}
	if b.s[0].Ypos < 0 {
		b.s[0].Ypos = 0
	}
	if b.s[0].Xpos > 9 {
		b.s[0].Xpos = 9
	}
	if b.s[0].Ypos > 9 {
		b.s[0].Ypos = 9
	}*/
}

func Init(numCols int, width int) {
	colWidth = width
	cols = numCols
}

func Create() Body {
	b := Body{}
	b.s = make([]cell.Cell, 1, cols * cols)
	b.Imd = imdraw.New(nil)
	b.direction = 0
	b.nextDir = 0
	b.addCell = false
	cell.Init(colWidth, b.Imd)
	b.s[0] = cell.Create(0, 0)
	for i := 1; i < 4; i++ {
		b.s = append(b.s, cell.Create(i, 0))
	}
	b.Redraw()
	return b
}