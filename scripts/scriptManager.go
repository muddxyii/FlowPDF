package scripts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

// ScriptType represents an enumerated type for defining different script action types for PDF operations.
type ScriptType int

const (
	PdfClear ScriptType = iota
	PdfMerge
)

type ScriptOptions struct {
	KeepInfo            bool
	KeepInitialTestData bool
	KeepRepairData      bool
	KeepFinalTestData   bool
}

func IsNodeInstalled() bool {
	cmd := exec.Command("node", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return err == nil
}

// RunScript executes a Node.js script for PDF operations such as clearing or merging PDFs based on the specified script type.
// The function takes a ScriptType, the URI of the PDF file, and an optional pointer to a PDF template string.
// It returns an error if the script type is invalid or if the command execution fails.
func RunScript(scriptType ScriptType, options *ScriptOptions, pdfURI string, pdfTemplate *string) error {
	action := ""
	switch scriptType {
	case PdfClear:
		action = "clearPDF"
	case PdfMerge:
		action = "mergePDF"
	default:
		return fmt.Errorf("invalid script type")
	}

	// Serialize ScriptOptions to JSON
	optionsJSON := "{}"
	if options != nil {
		optionsBytes, err := json.Marshal(options)
		if err != nil {
			return fmt.Errorf("failed to serialize ScriptOptions: %v", err)
		}
		optionsJSON = string(optionsBytes) // Serialize as string
	}

	templateArg := ""
	if pdfTemplate != nil {
		templateArg = *pdfTemplate
	}

	cmd := exec.Command("node", "scripts/pdf_scripts.js", action, pdfURI, templateArg, optionsJSON)
	return cmd.Run()
}
