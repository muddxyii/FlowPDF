const { PDFDocument } = require('pdf-lib');
const fs = require('fs');

async function clearPDF(pdfPath) {
    const pdfBytes = await fs.promises.readFile(pdfPath);
    const pdfDoc = await PDFDocument.load(pdfBytes);

    const form = pdfDoc.getForm();
    const serialNoField = form.getTextField('SerialNo');
    serialNoField.setText('');

    const clearedPdf = await pdfDoc.save();
    await fs.promises.writeFile(pdfPath, clearedPdf);
}

async function mergePDF(pdfPath) {
    console.log(`Merging PDFs for: ${pdfPath}`);
    // Logic to merge PDFs
}


// Handle CLI argument
const action = process.argv[2];
const pdfPath = process.argv[3];

if (pdfPath && action === 'clearPDF') {
    clearPDF(pdfPath).catch(console.error);
}
else if (pdfPath && action === 'mergePDF') {
    const templatePath = process.argv[4];
    mergePDF(pdfPath).catch(console.error);
}