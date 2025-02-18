package main

import (
	"FlowPDF/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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
	switchPage(ui.MainContentPage())

	copyrightLabel := widget.NewLabel("2025 © AnyBackflow.com Inc.")
	versionLabel := widget.NewLabel("FlowPDF v1.0.0")
	versionContainerWithSeparator := container.NewVBox(
		widget.NewSeparator(),
		container.NewHBox(
			layout.NewSpacer(),
			copyrightLabel,
			widget.NewSeparator(),
			versionLabel,
		),
	)

	// Layout: Navigation on the left, dynamic content on the right
	content := container.NewBorder(nil, versionContainerWithSeparator, navigationBar, nil, mainContent)
	window.SetContent(container.NewPadded(content))

	window.ShowAndRun()
}
