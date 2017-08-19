package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	//"github.com/faiface/pixel/imdraw"
	"github.com/cdalizadeh/snake/field"
	"github.com/cdalizadeh/snake/body"
	"github.com/cdalizadeh/snake/food"
	"math/rand"
	//"fmt"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	var width int = 800
	var cols int = 10
	var colWidth int = int(float64(width) / float64(cols))
	bColor := pixel.RGB(0, 0, 0)
	var timerConstant int = 8
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
	body.Init(cols, colWidth)
	food.Init(colWidth)
	snakeBody := body.Create()
	snakeFood := food.Create(rand.Intn(cols), rand.Intn(cols))
	snakeField := field.Create(width, cols, colWidth)

	for !win.Closed() {
		timer--
		if timer <= 0 {
			timer = timerConstant
			snakeBody.Move()
			headx, heady := snakeBody.GetHead()
			if snakeFood.Xpos == headx && snakeFood.Ypos == heady {
				snakeBody.Eat()
				foodx := rand.Intn(cols)
				foody := rand.Intn(cols)
				for snakeBody.IsWithinBody(foodx, foody) {
					foodx = rand.Intn(cols)
					foody = rand.Intn(cols)
				}
				snakeFood.Set(foodx, foody)
			}
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
		snakeFood.Imd.Draw(win)
		snakeBody.Imd.Draw(win)
		snakeField.Imd.Draw(win)
		win.Update()
	}
}