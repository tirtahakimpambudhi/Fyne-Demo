package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"go_fyne_demo/config"
)

func main(){
	// instance app class to object
	application := app.New()
	// instance window class to object
	//  params title string 
	window := application.NewWindow("Hyper Link")
	// resize window
	// params fyne.size struct
	window.Resize(fyne.Size{Width: config.Width, Height: config.Height})
}