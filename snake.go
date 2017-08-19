package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	//"github.com/faiface/pixel/imdraw"
	"github.com/cdalizadeh/snake/field"
	"github.com/cdalizadeh/snake/body"
	//"github.com/cdalizadeh/snake/cell"
	//"fmt"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	var width int = 900
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
	body.Init(cols, colWidth)
	snakeBody := body.Create()
	for !win.Closed() {
		timer--
		if timer <= 0 {
			timer = timerConstant
			snakeBody.Move()
		}
		if win.JustPressed(pixelgl.KeyLeft) {
			snakeBody.SetDir(2)
		}
		if win.JustPressed(pixelgl.KeyRight) {
			snakeBody.SetDir(0)
		}
		if win.JustPressed(pixelgl.KeyDown) {
			snakeBody.SetDir(3)
		}
		if win.JustPressed(pixelgl.KeyUp) {
			snakeBody.SetDir(1)
		}
		win.Clear(bColor)
		snakeBody.Imd.Draw(win)
		fieldImd.Draw(win)
		win.Update()
	}
}