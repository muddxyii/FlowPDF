package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BuildFooterContent() fyne.CanvasObject {
	copyrightLabel := widget.NewLabel("2025 © AnyBackflow.com Inc.")
	versionLabel := widget.NewLabel("FlowPDF v1.0.0")
	footer := container.NewVBox(
		widget.NewSeparator(),
		container.NewHBox(
			layout.NewSpacer(),
			copyrightLabel,
			widget.NewSeparator(),
			versionLabel,
		),
	)

	return footer
}
