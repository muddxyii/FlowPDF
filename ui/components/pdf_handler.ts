import { PDFDocument } from 'pdf-lib'
import * as fs from "node:fs";

async function clearForms(pdfPath: string): Promise<void> {
    const pdfBytes = await fs.promises.readFile(pdfPath)
    const pdfDoc = await PDFDocument.load(pdfBytes)

    const form = pdfDoc.getForm()
    form.flatten()

    const clearedPdf = await pdfDoc.save()
    await fs.promises.writeFile(pdfPath, clearedPdf)
}

// Handle CLI argument
const pdfPath = process.argv[2]
if (pdfPath) {
    clearForms(pdfPath).catch(console.error)
}