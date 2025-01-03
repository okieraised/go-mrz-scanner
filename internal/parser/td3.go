package parser

import (
	"github.com/okieraised/go-mrz-scanner/constants"
	"github.com/okieraised/go-mrz-scanner/mrz_errors"
	"github.com/okieraised/go-mrz-scanner/utils"
	"strings"
)

type TD3 struct {
}

func NewTD3() IMRZParser {
	return &TD3{}
}

func (td3 *TD3) Parse(in []string) (*MRZResult, error) {

	result := &MRZResult{}
	parsedResult := make(map[string]*mrzField)

	if len(in) != 2 {
		return result, mrz_errors.ErrGenericInvalidMRZLength
	}

	for _, line := range in {
		if len(line) != constants.Type3NumberOfCharactersPerLine {
			return result, mrz_errors.ErrTD3InvalidLineLength
		}
	}

	isVisaDocument := false
	firstLine, secondLine := in[0], in[1]
	formatter := NewFieldFormatter(true)

	if firstLine[0] == 'V' {
		isVisaDocument = true
	}
	result.IsVISA = isVisaDocument

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

	name, err := formatter.field(namesField, firstLine, 5, 39, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldName] = name

	// parse second line
	documentNumber, err := formatter.field(documentNumberField, secondLine, 0, 9, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldDocumentNumber] = documentNumber

	nationality, err := formatter.field(nationalityField, secondLine, 10, 3, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldNationality] = nationality

	birthdate, err := formatter.field(birthdateField, secondLine, 13, 6, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldBirthdate] = birthdate

	sex, err := formatter.field(sexField, secondLine, 20, 1, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldGender] = sex

	expiryDate, err := formatter.field(expiryDateField, secondLine, 21, 6, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.MRZFieldExpiryDate] = expiryDate

	var personalNumber *mrzField = nil
	var finalCheckDigit *mrzField = nil

	if isVisaDocument {
		personalNumber, err = formatter.field(personalNumberField, secondLine, 28, 16, false)
		if err != nil {
			return result, err
		}
		parsedResult[constants.MRZFieldOptionalData1] = personalNumber
		parsedResult[constants.MRZFieldFinalCheckDigit] = nil
	} else {
		personalNumber, err = formatter.field(personalNumberField, secondLine, 28, 14, true)
		if err != nil {
			return result, err
		}
		parsedResult[constants.MRZFieldOptionalData1] = personalNumber

		finalCheckDigit, err = formatter.field(hashField, secondLine, 43, 1, false)
		if err != nil {
			return result, err
		}
		parsedResult[constants.MRZFieldFinalCheckDigit] = finalCheckDigit
	}

	isValid, err := td3.validateAllCheckDigits(documentNumber, birthdate, expiryDate, personalNumber, finalCheckDigit)
	if err != nil {
		return result, err
	}

	result.Fields = parsedResult
	result.IsValid = isValid
	result.IssuingState = constants.IssuingCountryCodes[countryCode.value.(string)]

	return result, nil
}

func (td3 *TD3) validateAllCheckDigits(documentNumber, birthdate, expiryDate, optionalData, finalCheckDigit *mrzField) (bool, error) {

	if finalCheckDigit != nil {
		compositeStr := strings.Join([]string{
			documentNumber.rawValue, documentNumber.checkDigit,
			birthdate.rawValue, birthdate.checkDigit,
			expiryDate.rawValue, expiryDate.checkDigit,
			optionalData.rawValue, optionalData.checkDigit,
		}, "")

		calculatedCheckDigit, err := utils.CalculateCheckDigits(compositeStr)
		if err != nil {
			return false, err
		}
		return documentNumber.isValid &&
			birthdate.isValid &&
			expiryDate.isValid &&
			optionalData.isValid &&
			(calculatedCheckDigit == finalCheckDigit.rawValue), nil
	} else {
		return documentNumber.isValid && birthdate.isValid && expiryDate.isValid, nil
	}
}
