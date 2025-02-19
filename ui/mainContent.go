package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainContentPage() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("FlowPDF - Manage Your PDF Forms", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Info section
	infoLabel := widget.NewLabelWithStyle(
		"Features:",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	clearPdfInfo := widget.NewLabelWithStyle(
		"• Clear PDF: \n"+
			"    - Removes all form entries in the currently loaded PDF.",
		fyne.TextAlignLeading,
		fyne.TextStyle{},
	)
	updateTemplateInfo := widget.NewLabelWithStyle(
		"• Update Template: \n"+
			"    - Converts old PDF formats to a newer version with updated dropdowns and reduced file size.",
		fyne.TextAlignLeading,
		fyne.TextStyle{},
	)

	// Instruction section
	instructionLabel := widget.NewLabelWithStyle(
		"Follow these steps to get started:",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)
	step1 := widget.NewLabel("1. Select the feature you want to use from the navigation bar.")
	step2 := widget.NewLabel("2. Attach the PDF form you want to use and then follow the on-screen instructions.")

	return container.NewVBox(
		title,
		widget.NewSeparator(),
		infoLabel,
		clearPdfInfo,
		updateTemplateInfo,
		widget.NewSeparator(),
		instructionLabel,
		step1,
		step2,
	)
}
