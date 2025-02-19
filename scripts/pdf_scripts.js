const { PDFDocument } = require('pdf-lib');
const fs = require('fs');

function getInfoFieldNames() {
    return ['WaterPurveyor'] +
        // Facility Owner
        ['FacilityOwner', 'Address', 'Email', 'Contact', 'Phone'] +
        // Representative Owner
        ['OwnerRep', 'RepAddress', 'PersontoContact', 'Phone-0'] +
        // Location Info
        ['AssemblyAddress', 'On Site Location of Assembly', 'PrimaryBusinessService'] +
        // Installation Info
        ['InstallationIs', 'ProtectionType', 'ServiceType'] +
        // Device Info 1
        ['SerialNo', 'WaterMeterNo', 'Size', 'ModelNo', 'SOVComment'] +
        // Device Info 2
        ['BFType', 'Manufacturer', 'SOVList'];
}

function getInitialFieldNames() {
    return ['DateFailed', 'InitialTester', 'InitialTesterNo', 'InitialTestKitSerial'] +
        ['LinePressure', 'InitialCT1', 'InitialCT2',
            'InitialPSIRV', 'InitialAirInlet', 'InitialCk1PVB'] +
        ['InitialCTBox', 'InitialCT1Leaked', 'InitialCT2Box', 'InitialCT2Leaked',
            'InitialRVDidNotOpen', 'InitialAirInletLeaked', 'InitialCkPVBLDidNotOpen',
            'InitialCkPVBLeaked'];
}

function getRepairsFieldNames() {
    return ['RepairedTester', 'RepairedTesterNo', 'DateRepaired', 'RepairedTestKitSerial'] +
        // Check Valve
        ['Ck1Cleaned', 'Ck1CheckDisc', 'Ck1DiscHolder',
            'Ck1Spring', 'Ck1Guide', 'Ck1Seat', 'Ck1Other', 'Ck2Cleaned', 'Ck2CheckDisc', 'Ck2DiscHolder',
            'Ck2Spring', 'Ck2Guide', 'Ck2Seat', 'Ck2Other',] +
        // Relief Valve
        ['RVCleaned', 'RVRubberKit', 'RVDiscHolder',
            'RVSpring', 'RVGuide', 'RVSeat', 'RVOther',] +
        // Vacuum Breaker
        ['PVBCleaned', 'PVBRubberKit', 'PVBDiscHolder',
            'PVBSpring', 'PVBGuide', 'PVBSeat', 'PVBOther'];
}

function getFinalFieldNames() {
    return ['DatePassed', 'FinalTester', 'FinalTesterNo', 'FinalTestKitSerial'] +
        ['LinePressure', 'FinalCT1', 'FinalCT2', 'FinalRV', 'FinalAirInlet', 'Check Valve'] +
        ['FinalCT1Box', 'FinalCT2Box', 'BackPressure']
}

function parseOptions(optionsString) {
    try {
        return JSON.parse(optionsString);
    } catch (err) {
        console.error('Failed to parse options:', err);
        return {};
    }
}

async function clearPDF(pdfPath, options) {
    const pdfBytes = await fs.promises.readFile(pdfPath);
    const pdfDoc = await PDFDocument.load(pdfBytes);

    const form = pdfDoc.getForm();
    const infoFieldsNames = getInfoFieldNames();
    const initialFieldsNames = getInitialFieldNames();
    const repairsFieldsNames = getRepairsFieldNames();
    const finalFieldsNames = getFinalFieldNames();

    const fields = form.getFields();
    fields.forEach(field => {
        if (options.KeepInfo && infoFieldsNames.includes(field.getName())) {
            return;
        }
        if (options.KeepComments && field.getName() === 'ReportComments') {
            return
        }
        if (options.KeepInitialTestData && initialFieldsNames.includes(field.getName())) {
            return;
        }
        if (options.KeepRepairData && repairsFieldsNames.includes(field.getName())) {
            return;
        }
        if (options.KeepFinalTestData && finalFieldsNames.includes(field.getName())) {
            return;
        }

        const fieldType = field.constructor.name;
        if (fieldType === 'PDFTextField') {
            field.setText('');
        } else if (fieldType === 'PDFCheckBox') {
            field.uncheck();
        } else if (fieldType === 'PDFRadioGroup') {
            field.clear();
        } else if (fieldType === 'PDFDropdown' || fieldType === 'PDFOptionList') {
            field.clear();
        }
    });

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
const templatePath = process.argv[4];
const optionsString = process.argv[5];

const options = parseOptions(optionsString);

if (pdfPath && action === 'clearPDF') {
    clearPDF(pdfPath, options).catch(console.error);
}
else if (pdfPath && templatePath && action === 'mergePDF') {
    mergePDF(pdfPath).catch(console.error);
}