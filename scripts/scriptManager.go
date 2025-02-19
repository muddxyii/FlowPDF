package scripts

import (
	"bytes"
	"os/exec"
)

type ScriptType int

const (
	PdfClear ScriptType = iota
	PdfMerge
)

func IsNodeInstalled() bool {
	cmd := exec.Command("node", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return err == nil
}

func RunScript(scriptType ScriptType, pdfURI string) error {
	switch scriptType {
	case PdfClear:
		cmd := exec.Command("node", "scripts/pdf_clear.js", pdfURI)
		return cmd.Run()
	case PdfMerge:
		break
	default:
		break
	}
	return nil
}
