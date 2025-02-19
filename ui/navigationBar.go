package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func BuildNavBar(switchPage func(newContent fyne.CanvasObject), win fyne.Window) *fyne.Container {
	// Home Button
	homeButton := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		switchPage(MainContentPage())
	})
	// Clear Button
	clearButton := widget.NewButtonWithIcon("Clear PDF", theme.ContentClearIcon(), func() {
		switchPage(ClearPdfContentPage(win))
	})
	// Update Format Button
	updateFormatButton := widget.NewButtonWithIcon("Update Template", theme.DocumentIcon(), func() {
		switchPage(UpdateTemplateContentPage(win))
	})
	// Edit Dropdown Button
	editDropdownButton := widget.NewButtonWithIcon("Edit Test IDs", theme.MoreVerticalIcon(), func() {
		switchPage(widget.NewLabel("Feature coming soon!"))
	})
	// Expand Button
	labelsHidden := false

	expandButton := widget.NewButtonWithIcon("Collapse", theme.ViewFullScreenIcon(), nil)
	expandButton.OnTapped = func() {
		if labelsHidden {
			homeButton.SetText("Home")
			clearButton.SetText("Clear PDF")
			updateFormatButton.SetText("Update Template")
			editDropdownButton.SetText("Edit Testers")
			expandButton.SetText("Collapse")
			expandButton.SetIcon(theme.ViewFullScreenIcon())
		} else {
			homeButton.SetText("")
			clearButton.SetText("")
			updateFormatButton.SetText("")
			editDropdownButton.SetText("")
			expandButton.SetText("")
			expandButton.SetIcon(theme.ViewRestoreIcon())
		}

		labelsHidden = !labelsHidden
	}

	buttons := container.NewVBox(
		container.NewPadded(homeButton),
		container.NewPadded(clearButton),
		container.NewPadded(updateFormatButton),
		container.NewPadded(editDropdownButton),
		layout.NewSpacer(),
		container.NewPadded(expandButton),
	)

	return container.NewHBox(buttons, widget.NewSeparator())
}
