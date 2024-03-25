package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go_fyne_demo/config"
)

func main() {
	// For instance class App
	application := app.New()
	// For instance class Window
	window := application.NewWindow("hello world")
	// Resizing width and height window
	window.Resize(fyne.NewSize(config.Width, config.Height))
	// Set content window with label
	window.SetContent(container.NewVBox(
		widget.NewLabel("Hello World!!"),
		widget.NewLabel("First GUI With Fyne Golang"),
	))
	// Logging
	fmt.Println("Running Application")
	// Showing window
	window.ShowAndRun()
}
