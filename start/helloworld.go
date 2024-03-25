package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const (
	width  float32 = 600
	height float32 = 400
)

func main() {
	application := app.New()
	window := application.NewWindow("hello world")
	window.Resize(fyne.Size{Width: width, Height: height})

	window.SetContent(widget.NewLabel("Hello World!!\nFirst GUI With Fyne Golang"))
	window.ShowAndRun()
}
