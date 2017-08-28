package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"math"
	"math/rand"
	"fmt"
)

//consts
var width float64 = 800
var cols int = 10
var colWidth float64 = width / float64(cols)
var backColor pixel.RGBA = pixel.RGB(0, 0, 0)
var lineColor pixel.RGBA = pixel.RGB(0, 1, 0)
var bodyColor pixel.RGBA = pixel.RGB(1, 1, 1)
var foodColor pixel.RGBA = pixel.RGB(0, 1, 1)
var timerConstant int = 11

//globals
var win *pixelgl.Window
var err error
var snakeBody Body
var snakeFood Food
var snakeField Field
var basicTxt *text.Text
var gameover bool = false

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Snake",
		Bounds: pixel.R(0, 0, width, width),
		VSync:  true,
	}
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt = text.New(pixel.V(100, width / 2), basicAtlas)
	basicTxt.Clear()
	basicTxt.Dot = basicTxt.Orig
	fmt.Fprintln(basicTxt, "PAUSE")

	for !win.Closed() {
		snakeBody = createBody(0, 0, 0, bodyColor)
		snakeFood = createFood(rand.Intn(cols), rand.Intn(cols), foodColor)
		snakeField = createField(width, cols, colWidth, lineColor)

		startMenu()

		timer := timerConstant
		gameover = false
		for !win.Closed() && !gameover {
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
				pauseMenu()
			}
			win.Clear(backColor)
			snakeFood.Imd.Draw(win)
			snakeBody.Imd.Draw(win)
			snakeField.Imd.Draw(win)
			win.Update()
		}
	}
}

func startMenu () {
	win.Clear(backColor)
	colorMask(0.5)
	snakeFood.Draw()
	snakeBody.Draw()
	snakeField.Draw()
	snakeFood.Imd.Draw(win)
	snakeBody.Imd.Draw(win)
	snakeField.Imd.Draw(win)
	basicTxt.Clear()
	basicTxt.Dot = basicTxt.Orig
	fmt.Fprintln(basicTxt, "BEGIN")
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, math.Floor(width/100)))
	for !win.JustPressed(pixelgl.KeyP) && !win.Closed() {
		win.Update()
	}
	colorMask(1)
	snakeFood.Draw()
	snakeBody.Draw()
	snakeField.Draw()
	win.Update()
}

func pauseMenu() {
	win.Update()
	win.Clear(backColor)
	colorMask(0.5)
	snakeFood.Draw()
	snakeBody.Draw()
	snakeField.Draw()
	snakeFood.Imd.Draw(win)
	snakeBody.Imd.Draw(win)
	snakeField.Imd.Draw(win)
	basicTxt.Clear()
	basicTxt.Dot = basicTxt.Orig
	fmt.Fprintln(basicTxt, "PAUSE")
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, math.Floor(width/100)))
	for !win.JustPressed(pixelgl.KeyP) && !win.Closed() {
		win.Update()
	}
	colorMask(1)
	snakeFood.Draw()
	snakeBody.Draw()
	snakeField.Draw()
}

func colorMask(m float64) {
	snakeFood.Imd.SetColorMask(pixel.RGB(m, m, m))
	snakeBody.Imd.SetColorMask(pixel.RGB(m, m, m))
	snakeField.Imd.SetColorMask(pixel.RGB(m, m, m))
}