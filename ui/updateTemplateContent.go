package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UpdateTemplateContentPage() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("FlowPDF - Update your PDF Template", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	attachButton := widget.NewButton("Attach Old PDF Form", func() {
		// TODO: Implement attach PDF form logic
	})

	dropdownLabel := widget.NewLabelWithStyle("Select a Template:", fyne.TextAlignLeading, fyne.TextStyle{Italic: true})
	dropdownOptions := []string{"Template 1", "Template 2"}
	dropdown := widget.NewSelect(dropdownOptions, func(selected string) {
		// Placeholder action on dropdown selection
	})

	updateButton := widget.NewButton("Update PDF", func() {
		// TODO: Implement merge PDF logic
	})

	return container.NewVBox(
		title,
		widget.NewSeparator(),
		attachButton,
		container.NewVBox(
			dropdownLabel,
			container.NewPadded(dropdown),
		),
		updateButton,
	)
}
