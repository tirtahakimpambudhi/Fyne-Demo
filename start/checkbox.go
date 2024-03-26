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
	// Variable to hold food selection
	foods := make(map[string]bool)
	//Instance class App to Object
	application := app.New()
	//Instance class Window to Object
	window := application.NewWindow("Checkbox")
	// Resize window size width and height
	window.Resize(fyne.Size{Width: config.Width, Height: config.Height})
	// Set Content
	foodLabel := widget.NewLabel("Choose Favorite Food")
	kebabCheckBox := widget.NewCheck("Kebab", func(b bool) {
		if b && !foods["Kebab"] {
			foods["Kebab"] = b
		}
	})
	burgerCheckBox := widget.NewCheck("Burger", func(b bool) {
		if b && !foods["Burger"] {
			foods["Burger"] = b
		}
	})
	window.SetContent(container.NewVBox(
		foodLabel,
		kebabCheckBox,
		burgerCheckBox,
	))

	fmt.Println("Running Application")
	// Running and Showing
	window.ShowAndRun()
	fmt.Println(foods)
}
