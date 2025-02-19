package ui

import (
	"FlowPDF/ui/components"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ClearPdfContentPage(win fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("FlowPDF - Clear PDF Forms", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	pdfSelector := components.CreatePDFSelector(win, func(fileURI string) {
		//pdfURI = fileURI
	})

	clearButton := widget.NewButton("Clear Forms", func() {
		// TODO: Implement clear PDF logic
	})

	return container.NewVBox(
		title,
		widget.NewSeparator(),
		pdfSelector,
		clearButton,
	)
}
