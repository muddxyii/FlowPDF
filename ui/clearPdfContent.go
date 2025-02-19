package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

func ClearPdfContentPage(win fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("FlowPDF - Clear PDF Forms", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	selectedPDFLabel := widget.NewLabel("No PDF Selected")

	attachButton := widget.NewButton("Attach PDF Form", func() {
		dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
			if err != nil || file == nil {
				return
			}

			// Validate the file as a PDF
			if file.URI().Extension() != ".pdf" {
				dialog.ShowInformation("Invalid File", "Please select a valid PDF file.", win)
				return
			}

			// Handle the PDF file (e.g., store the path or process content)
			selectedPDFLabel.SetText("Selected PDF: " + file.URI().Name())
			log.Printf("Selected PDF: %s\n", file.URI().Path())
			defer func(file fyne.URIReadCloser) {
				err := file.Close()
				if err != nil {
					return
				}
			}(file)

		}, win)
	})

	clearButton := widget.NewButton("Clear Forms", func() {
		// TODO: Implement clear PDF logic
	})

	return container.NewVBox(
		title,
		widget.NewSeparator(),
		attachButton,
		selectedPDFLabel,
		clearButton,
	)
}
