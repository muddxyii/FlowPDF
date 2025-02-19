const { PDFDocument } = require('pdf-lib');
const fs = require('fs');

async function clearForms(pdfPath) {
    const pdfBytes = await fs.promises.readFile(pdfPath);
    const pdfDoc = await PDFDocument.load(pdfBytes);

    const form = pdfDoc.getForm();
    const serialNoField = form.getTextField('SerialNo');
    serialNoField.setText('');

    const clearedPdf = await pdfDoc.save();
    await fs.promises.writeFile(pdfPath, clearedPdf);
}

// Handle CLI argument
const pdfPath = process.argv[2];
if (pdfPath) {
    clearForms(pdfPath).catch(console.error);
}