package ui

import (
	"FlowPDF/scripts"
	"FlowPDF/ui/components"
	"embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/pkg/browser"
	"path/filepath"
	"strings"
)

//go:embed assets/templates/*.pdf
var embeddedTemplates embed.FS

func getTemplateFiles() ([]string, error) {
	var templates []string

	entries, err := embeddedTemplates.ReadDir("assets/templates")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded templates: %w", err)
	}

	// Add files with .pdf extension
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".pdf") {
			templates = append(templates, entry.Name())
		}
	}

	return templates, nil
}

func UpdateTemplateContentPage(win fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("FlowPDF - Update your PDF Template", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	var pdfURI string
	var selectedTemplateURI string

	pdfSelector := components.CreatePDFSelector(win, func(fileURI string) {
		pdfURI = fileURI
	})

	dropdownOptions, err := getTemplateFiles()
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to get template files: %v", err), win)
		dropdownOptions = []string{}
	}
	dropdownLabel := widget.NewLabelWithStyle("Select a Template:", fyne.TextAlignLeading, fyne.TextStyle{Italic: true})
	dropdown := widget.NewSelect(dropdownOptions, func(selected string) {
		selectedTemplateURI = filepath.Join("assets/templates", selected)
		fmt.Printf("Template selected: %s\n", selectedTemplateURI)
	})
	if len(dropdownOptions) > 0 {
		dropdown.SetSelected(dropdownOptions[0])
		selectedTemplateURI = filepath.Join("assets/templates", dropdownOptions[0])
		fmt.Printf("Default template selected: %s\n", selectedTemplateURI)
	}

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
		if selectedTemplateURI == "" {
			dialog.ShowError(fmt.Errorf("please select a template"), win)
			return
		}

		err := scripts.RunScript(scripts.PdfMerge, nil, pdfURI, &selectedTemplateURI)
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
