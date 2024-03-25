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
	// Variable for trigger button every click increment
	var like int
	// For instance class App
	application := app.New()
	// For instance class Window
	window := application.NewWindow("Button")
	// Resizing width and height window
	window.Resize(fyne.NewSize(config.Width, config.Height))
	// Set content window with label
	window.SetContent(container.NewVBox(
		widget.NewLabel("Try Button in Fyne Golang!!"),
		// Set New Button
		widget.NewButton("click!!", func() {
			// Trigger every button clicked
			like++
			fmt.Println(like)
		}),
	))
	// Logging
	fmt.Println("Running Application")
	// Showing window
	window.ShowAndRun()
}
