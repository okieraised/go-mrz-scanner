package parser

import (
	"github.com/okieraised/go-mrz-scanner/constants"
	"github.com/okieraised/go-mrz-scanner/mrz_errors"
	"github.com/okieraised/go-mrz-scanner/utils"
	"strings"
)

type TD2 struct {
}

func NewTD2() IMRZParser {
	return &TD2{}
}

func (td2 *TD2) Parse(in []string) (*ParserResult, error) {

	result := &ParserResult{}
	parsedResult := make(map[string]*mrzField)

	if len(in) != 2 {
		return result, mrz_errors.ErrGenericInvalidMRZLength
	}

	for _, line := range in {
		if len(line) != constants.Type2NumberOfCharacter {
			return result, mrz_errors.ErrTD2InvalidLineLength
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
	parsedResult[constants.DocumentType] = documentType

	countryCode, err := formatter.field(countryCodeField, firstLine, 2, 3, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.CountryCode] = countryCode

	name, err := formatter.field(namesField, firstLine, 5, 31, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.Name] = name

	// parse second line
	documentNumber, err := formatter.field(documentNumberField, secondLine, 0, 9, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.DocumentNumber] = documentNumber

	nationality, err := formatter.field(nationalityField, secondLine, 10, 3, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.Nationality] = nationality

	birthdate, err := formatter.field(birthdateField, secondLine, 13, 6, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.Birthdate] = birthdate

	sex, err := formatter.field(sexField, secondLine, 20, 1, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.Sex] = sex

	expiryDate, err := formatter.field(expiryDateField, secondLine, 21, 6, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.ExpiryDate] = expiryDate

	var optionalData *mrzField = nil
	var finalCheckDigit *mrzField = nil

	if isVisaDocument {
		optionalData, err = formatter.field(personalNumberField, secondLine, 28, 8, false)
		if err != nil {
			return result, err
		}
		parsedResult[constants.OptionalData1] = optionalData

		parsedResult[constants.FinalCheckDigit] = nil
	} else {
		optionalData, err = formatter.field(personalNumberField, secondLine, 28, 7, true)
		if err != nil {
			return result, err
		}
		parsedResult[constants.OptionalData1] = optionalData

		finalCheckDigit, err = formatter.field(hashField, secondLine, 35, 1, false)
		if err != nil {
			return result, err
		}
		parsedResult[constants.FinalCheckDigit] = finalCheckDigit
	}

	isValid, err := td2.validateAllCheckDigits(documentNumber, birthdate, expiryDate, optionalData, finalCheckDigit)
	if err != nil {
		return result, err
	}

	result.Mapper = parsedResult
	result.IsValid = isValid

	return result, nil
}

func (td2 *TD2) validateAllCheckDigits(documentNumber, birthdate, expiryDate, optionalData, finalCheckDigit *mrzField) (bool, error) {

	if finalCheckDigit != nil {
		compositeStr := strings.Join([]string{
			documentNumber.rawValue, documentNumber.checkDigit,
			birthdate.rawValue, birthdate.checkDigit,
			expiryDate.rawValue, expiryDate.checkDigit,
			optionalData.rawValue,
		}, "")

		calculatedCheckDigit, err := utils.CalculateCheckDigits(compositeStr)
		if err != nil {
			return false, err
		}

		return documentNumber.isValid &&
			birthdate.isValid &&
			expiryDate.isValid &&
			(calculatedCheckDigit == finalCheckDigit.rawValue), nil
	} else {
		return documentNumber.isValid && birthdate.isValid && expiryDate.isValid, nil
	}
}