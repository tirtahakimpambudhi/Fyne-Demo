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
	var count int
	// Variable to contains count
	likeLabel := widget.NewLabel("")
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
			count++
			updateLikeLabel(likeLabel, count) // Update likeLabel every time the button is clicked
		}),
		likeLabel,
	))
	// Logging
	fmt.Println("Running Application")
	// Showing window
	window.ShowAndRun()
}

func updateLikeLabel(likeLabel *widget.Label, count int) {
	likeLabel.SetText(fmt.Sprintf("Like : %d", count))
}
