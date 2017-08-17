package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	//"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"github.com/cdalizadeh/snake/field"
	//"fmt"
	//"reflect"
)

func run() {
	var width int = 800
	var cols int = 20
	var colWidth int = int(float64(width) / float64(cols))
	
	cfg := pixelgl.WindowConfig{
		Title:  "Snake",
		Bounds: pixel.R(0, 0, float64(width), float64(width)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	field := field.Create(width, cols, colWidth)

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		field.Draw(win)
		win.Update()
	}

}

func main() {
	pixelgl.Run(run)
}
