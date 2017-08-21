package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"github.com/cdalizadeh/snake/field"
	"github.com/cdalizadeh/snake/body"
	"github.com/cdalizadeh/snake/food"
	"math/rand"
	"fmt"
)

var width float64 = 800
var cols int = 10
var colWidth float64 = width / float64(cols)
var backColor pixel.RGBA = pixel.RGB(0, 0, 0)
var lineColor pixel.RGBA = pixel.RGB(0, 1, 0)
var bodyColor pixel.RGBA = pixel.RGB(1, 1, 1)
var foodColor pixel.RGBA = pixel.RGB(0, 1, 1)
var timerConstant int = 11

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Snake",
		Bounds: pixel.R(0, 0, width, width),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(100, width / 2), basicAtlas)
	fmt.Fprintln(basicTxt, "PAUSE")
	
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
			win.Clear(backColor)
			snakeFood.Imd.SetColorMask(pixel.RGB(0.5, 0.5, 0.5))
			snakeBody.Imd.SetColorMask(pixel.RGB(0.5, 0.5, 0.5))
			snakeField.Imd.SetColorMask(pixel.RGB(0.5, 0.5, 0.5))
			snakeFood.Draw()
			snakeBody.Draw()
			snakeField.Draw()
			snakeFood.Imd.Draw(win)
			snakeBody.Imd.Draw(win)
			snakeField.Imd.Draw(win)
			basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 9))
			for !win.JustPressed(pixelgl.KeyP) && !win.Closed() {
				win.Update()
			}
			snakeFood.Imd.SetColorMask(pixel.RGB(1, 1, 1))
			snakeBody.Imd.SetColorMask(pixel.RGB(1, 1, 1))
			snakeField.Imd.SetColorMask(pixel.RGB(1, 1, 1))
			snakeFood.Draw()
			snakeBody.Draw()
			snakeField.Draw()
		}
		win.Clear(backColor)
		snakeFood.Imd.Draw(win)
		snakeBody.Imd.Draw(win)
		snakeField.Imd.Draw(win)
		win.Update()
	}
}