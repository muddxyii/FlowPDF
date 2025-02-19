package ui

import (
	"FlowPDF/scripts"
	"FlowPDF/ui/components"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/pkg/browser"
)

func UpdateTemplateContentPage(win fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("FlowPDF - Update your PDF Template", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	var pdfURI string

	pdfSelector := components.CreatePDFSelector(win, func(fileURI string) {
		pdfURI = fileURI
	})

	dropdownLabel := widget.NewLabelWithStyle("Select a Template:", fyne.TextAlignLeading, fyne.TextStyle{Italic: true})
	dropdownOptions := []string{"Template 1", "Template 2"}
	dropdown := widget.NewSelect(dropdownOptions, func(selected string) {
		// Placeholder action on dropdown selection
	})

	updateButton := widget.NewButton("Update PDF", func() {
		if !scripts.IsNodeInstalled() {
			dialog.NewConfirm(
				"Node.js Required",
				"Node.js is not installed. Please install Node.js to use this feature.\n"+
					"Would you like to visit the download page?",
				func(confirmed bool) {
					if confirmed {
						err := browser.OpenURL("https://nodejs.org/")
						if err != nil {
							panic(err)
						}
					}
				},
				win,
			).Show()
			return
		}

		if pdfURI == "" {
			dialog.ShowError(fmt.Errorf("please select a PDF file"), win)
			return
		}

		err := scripts.RunScript(scripts.PdfMerge, nil, pdfURI, nil)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to update PDF: %v", err), win)
			return
		}

		dialog.ShowInformation("Success", "PDF updated successfully!", win)
	})

	return container.NewVBox(
		title,
		widget.NewSeparator(),
		pdfSelector,
		container.NewVBox(
			dropdownLabel,
			container.NewPadded(dropdown),
		),
		updateButton,
	)
}
