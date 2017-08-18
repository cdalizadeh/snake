package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/cdalizadeh/snake/field"
	"github.com/cdalizadeh/snake/body"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	var width int = 800
	var cols int = 10
	var colWidth int = int(float64(width) / float64(cols))
	bColor := pixel.RGB(0, 0, 0)
	var timerConstant int = 10
	var timer int = timerConstant
	
	cfg := pixelgl.WindowConfig{
		Title:  "Snake",
		Bounds: pixel.R(0, 0, float64(width), float64(width)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(bColor)
	fieldImd := field.Create(width, cols, colWidth)
	bodyImd := body.Create(colWidth)
	for !win.Closed() {
		timer--
		if timer <= 0 {
			timer = timerConstant
			body.Move()
		}
		if win.JustPressed(pixelgl.KeyLeft) {
			body.SetDir(2)
		}
		if win.JustPressed(pixelgl.KeyRight) {
			body.SetDir(0)
		}
		if win.JustPressed(pixelgl.KeyDown) {
			body.SetDir(3)
		}
		if win.JustPressed(pixelgl.KeyUp) {
			body.SetDir(1)
		}
		win.Clear(bColor)
		bodyImd.Draw(win)
		fieldImd.Draw(win)
		win.Update()
	}
}