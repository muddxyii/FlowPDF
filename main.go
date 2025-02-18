package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("FlowPDF")

	label := widget.NewLabel("Hello Fyne!")
	window.SetContent(container.NewVBox(label))

	window.Resize(fyne.NewSize(200, 100))
	window.ShowAndRun()
}
