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

func ClearPdfContentPage(win fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("FlowPDF - Clear PDF Forms", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	var pdfURI string

	pdfSelector := components.CreatePDFSelector(win, func(fileURI string) {
		pdfURI = fileURI
	})

	var options scripts.ScriptOptions

	keepInfoCheck := widget.NewCheck("Keep Information", func(checked bool) {
		options.KeepInfo = checked
	})

	subCheckboxContainer := container.NewHBox()
	keepTestDataCheck := widget.NewCheck("Keep Test Data", func(checked bool) {
		if checked {
			initialTestDataCheck := widget.NewCheck("Initial Test Data", func(checked bool) {
				options.KeepInitialTestData = checked
			})
			initialTestDataCheck.SetChecked(true)

			repairDataCheck := widget.NewCheck("Repair Data", func(checked bool) {
				options.KeepRepairData = checked
			})
			repairDataCheck.SetChecked(false)

			finalTestDataCheck := widget.NewCheck("Final Test Data", func(checked bool) {
				options.KeepFinalTestData = checked
			})
			finalTestDataCheck.SetChecked(false)

			subCheckboxContainer.Objects = []fyne.CanvasObject{
				initialTestDataCheck,
				repairDataCheck,
				finalTestDataCheck,
			}
			subCheckboxContainer.Refresh()
		} else {
			options = scripts.ScriptOptions{
				KeepInfo:            false,
				KeepComments:        false,
				KeepInitialTestData: false,
				KeepRepairData:      false,
				KeepFinalTestData:   false,
			}
			subCheckboxContainer.Objects = nil
			subCheckboxContainer.Refresh()
		}
	})

	keepCommentsCheck := widget.NewCheck("Keep Comments", func(checked bool) {
		options.KeepComments = checked
	})

	clearButton := widget.NewButton("Clear Forms", func() {
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

		err := scripts.RunScript(scripts.PdfClear, &options, pdfURI, nil)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to clear forms: %v", err), win)
			return
		}

		dialog.ShowInformation("Success", "PDF forms cleared successfully!", win)
	})

	return container.NewVBox(
		title,
		widget.NewSeparator(),
		pdfSelector,
		container.NewHBox(keepInfoCheck, keepTestDataCheck, keepCommentsCheck),
		subCheckboxContainer,
		clearButton,
	)
}
