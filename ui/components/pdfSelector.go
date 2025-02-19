package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

// CreatePDFSelector returns a reusable UI component for selecting PDF files.
func CreatePDFSelector(win fyne.Window, onPDFSelected func(fileURI string)) fyne.CanvasObject {
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

			// Callback to handle the PDF file (e.g., set label, process file)
			selectedPDFLabel.SetText("Selected PDF: " + file.URI().Name())
			log.Printf("Selected PDF: %s\n", file.URI().Path())
			if onPDFSelected != nil {
				onPDFSelected(file.URI().Path())
			}

			defer func(file fyne.URIReadCloser) {
				err := file.Close()
				if err != nil {
					log.Printf("Error closing file: %v", err)
				}
			}(file)
		}, win)
	})

	return container.NewVBox(attachButton, selectedPDFLabel)
}
