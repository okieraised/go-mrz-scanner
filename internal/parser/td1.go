package parser

import (
	"github.com/okieraised/go-mrz-scanner/constants"
	"github.com/okieraised/go-mrz-scanner/mrz_errors"
	"github.com/okieraised/go-mrz-scanner/utils"
	"strings"
)

type TD1 struct {
}

func NewTD1() IMRZParser {
	return &TD1{}
}

func (td1 *TD1) Parse(in []string) (*MRZResult, error) {

	result := &MRZResult{
		IsVISA: false,
	}
	parsedResult := make(map[string]*mrzField)

	if len(in) != 3 {
		return result, mrz_errors.ErrGenericInvalidMRZLength
	}

	for _, line := range in {
		if len(line) != constants.Type1NumberOfCharactersPerLine {
			return result, mrz_errors.ErrTD1InvalidLineLength
		}
	}
	firstLine, secondLine, thirdLine := in[0], in[1], in[2]
	formatter := NewFieldFormatter(true)

	// parse first line
	documentType, err := formatter.field(documentTypeField, firstLine, 0, 2, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldDocumentType] = documentType

	countryCode, err := formatter.field(countryCodeField, firstLine, 2, 3, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldCountryCode] = countryCode

	documentNumber, err := formatter.field(documentNumberField, firstLine, 5, 9, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldDocumentNumber] = documentNumber

	optionalData1, err := formatter.field(optionalDataField, firstLine, 15, 15, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldOptionalData1] = optionalData1

	// Parse second line
	birthdate, err := formatter.field(birthdateField, secondLine, 0, 6, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldBirthdate] = birthdate

	sex, err := formatter.field(sexField, secondLine, 7, 1, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldGender] = sex

	expiryDate, err := formatter.field(expiryDateField, secondLine, 8, 6, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldExpiryDate] = expiryDate

	nationality, err := formatter.field(nationalityField, secondLine, 15, 3, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldNationality] = nationality

	optionalData2, err := formatter.field(optionalDataField, secondLine, 18, 11, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldOptionalData2] = optionalData2

	finalCheckDigit, err := formatter.field(hashField, secondLine, 29, 1, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldFinalCheckDigit] = finalCheckDigit

	// Parse third line
	name, err := formatter.field(namesField, thirdLine, 0, 30, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldName] = name

	isValid, err := td1.validateAllCheckDigits(documentNumber, optionalData1, birthdate, expiryDate, optionalData2, finalCheckDigit)
	if err != nil {
		return result, err
	}

	result.Fields = parsedResult
	result.IsValid = isValid

	result.IssuingState = constants.IssuingCountryCodes[countryCode.value.(string)]

	return result, nil
}

func (td1 *TD1) validateAllCheckDigits(documentNumber, optionalData, birthdate, expiryDate, optionalData2, finalCheckDigit *mrzField) (bool, error) {
	compositeStr := strings.Join([]string{
		documentNumber.rawValue, documentNumber.checkDigit,
		optionalData.rawValue, optionalData.checkDigit,
		birthdate.rawValue, birthdate.checkDigit,
		expiryDate.rawValue, expiryDate.checkDigit,
		optionalData2.rawValue, optionalData2.checkDigit,
	}, "")

	calculatedCheckDigit, err := utils.CalculateCheckDigits(compositeStr)
	if err != nil {
		return false, err
	}
	return documentNumber.isValid && birthdate.isValid && expiryDate.isValid && (calculatedCheckDigit == finalCheckDigit.value), nil
}
