package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainContentPage() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("FlowPDF - Manage Your PDF Forms", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	attachButton := widget.NewButton("Attach PDF Form", func() {
		// TODO: Implement attach PDF form logic
	})

	dropdownLabel := widget.NewLabelWithStyle("Edit Dropdowns:", fyne.TextAlignLeading, fyne.TextStyle{Italic: true})
	dropdownOptions := []string{"Option 1", "Option 2", "Option 3"}
	dropdown := widget.NewSelect(dropdownOptions, func(selected string) {
		// Placeholder action on dropdown selection
	})

	textInputLabel := widget.NewLabelWithStyle("Enter Text Input:", fyne.TextAlignLeading, fyne.TextStyle{})
	textInput := widget.NewEntry()
	textInput.SetPlaceHolder("Type here...")

	return container.NewVBox(
		title,
		widget.NewSeparator(),
		attachButton,
		container.NewVBox(
			dropdownLabel,
			container.NewPadded(dropdown),
		),
		container.NewVBox(
			textInputLabel,
			container.NewPadded(textInput),
		),
	)
}
