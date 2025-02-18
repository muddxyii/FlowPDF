package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ClearPdfContentPage() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("FlowPDF - Clear PDF Forms", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	attachButton := widget.NewButton("Attach PDF Form", func() {
		// TODO: Implement attach PDF form logic
	})

	clearButton := widget.NewButton("Clear PDF", func() {
		// TODO: Implement clear PDF logic
	})

	return container.NewVBox(
		title,
		widget.NewSeparator(),
		attachButton,
		clearButton,
	)
}
