package main

import (
	"FlowPDF/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("FlowPDF")
	window.Resize(fyne.NewSize(600, 400))

	mainContent := container.NewStack()

	// Helper function to switch pages
	switchPage := func(newContent fyne.CanvasObject) {
		mainContent.Objects = []fyne.CanvasObject{newContent}
		mainContent.Refresh()
	}

	// Build UI components
	navigationBar := ui.BuildNavBar(switchPage)
	switchPage(ui.MainContentPage()) // Set default home page

	// Layout: Navigation on the left, dynamic content on the right
	content := container.NewBorder(nil, nil, navigationBar, nil, mainContent)
	window.SetContent(container.NewPadded(content))

	window.ShowAndRun()
}
