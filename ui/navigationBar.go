package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func BuildNavBar(switchPage func(newContent fyne.CanvasObject)) *fyne.Container {
	// Home Button
	homeButton := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		switchPage(MainContentPage())
	})
	// Clear Button
	clearButton := widget.NewButtonWithIcon("Clear PDF", theme.ContentClearIcon(), func() {
		switchPage(ClearPdfContentPage())
	})
	// Update Format Button
	updateFormatButton := widget.NewButtonWithIcon("Update Template", theme.DocumentIcon(), func() {
		switchPage(UpdateTemplateContentPage())
	})
	// Edit Dropdown Button
	editDropdownButton := widget.NewButtonWithIcon("Edit Testers", theme.MoreVerticalIcon(), func() {
		switchPage(widget.NewLabel("Edit Dropdowns Page - Content goes here!"))
	})

	buttons := container.NewVBox(
		container.NewPadded(homeButton),
		container.NewPadded(clearButton),
		container.NewPadded(updateFormatButton),
		container.NewPadded(editDropdownButton),
	)

	return container.NewHBox(buttons, widget.NewSeparator())
}
