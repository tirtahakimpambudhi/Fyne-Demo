package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// constant width and height window
const (
	width  float32 = 400
	height float32 = 600
)

func main() {
	//for instance class App
	application := app.New()
	//for instance class Window
	window := application.NewWindow("hello world")
	//Resizing width and height window
	window.Resize(fyne.Size{Width: width, Height: height})
	// Set content window with label
	window.SetContent(widget.NewLabel("Hello World!!\nFirst GUI With Fyne Golang"))
	// showing window and running
	window.ShowAndRun()
}
