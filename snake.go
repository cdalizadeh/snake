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
	width := 800
	cols := 10
	colWidth := int(float64(width) / float64(cols))
	backColor := pixel.RGB(0, 0, 0)
	lineColor := pixel.RGB(0, 1, 0)
	bodyColor := pixel.RGB(1, 1, 1)
	foodColor := pixel.RGB(0, 1, 1)
	timerConstant := 11
	
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
	snakeBody := body.Create(0, 0, 0, bodyColor)
	snakeFood := food.Create(rand.Intn(cols), rand.Intn(cols), foodColor)
	snakeField := field.Create(width, cols, colWidth, lineColor)

	timer := timerConstant
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
		if win.JustPressed(pixelgl.KeyP) {
			win.Update()
			for !win.JustPressed(pixelgl.KeyP) && !win.Closed() {
				win.Update()
			}
		}
		win.Clear(backColor)
		snakeFood.Imd.Draw(win)
		snakeBody.Imd.Draw(win)
		snakeField.Imd.Draw(win)
		win.Update()
	}
}