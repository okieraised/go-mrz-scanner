package parser

import (
	"fmt"
	"github.com/okieraised/go-mrz-scanner/constants"
	"github.com/okieraised/go-mrz-scanner/mrz_errors"
	"github.com/okieraised/go-mrz-scanner/utils"
	"strings"
)

type TD1 struct {
	lineLength int
}

func NewTD1() IMRZParser {
	return &TD1{
		lineLength: 30,
	}
}

func (td1 *TD1) Parse(in []string) (*ParserResult, error) {

	result := &ParserResult{}
	parsedResult := make(map[string]*mrzField)

	if len(in) != 3 {
		return result, mrz_errors.ErrGenericInvalidMRZLength
	}

	for _, line := range in {
		if len(line) != 30 {
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
	parsedResult[constants.DocumentType] = documentType

	countryCode, err := formatter.field(countryCodeField, firstLine, 2, 3, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.CountryCode] = countryCode

	documentNumber, err := formatter.field(documentNumberField, firstLine, 5, 9, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.DocumentNumber] = documentNumber

	optionalData1, err := formatter.field(optionalDataField, firstLine, 15, 15, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.OptionalData1] = optionalData1

	// Parse second line
	birthdate, err := formatter.field(birthdateField, secondLine, 0, 6, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.Birthdate] = birthdate

	sex, err := formatter.field(sexField, secondLine, 7, 1, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.Sex] = sex

	expiryDate, err := formatter.field(expiryDateField, secondLine, 8, 6, true)
	if err != nil {
		return result, err
	}
	parsedResult[constants.ExpiryDate] = expiryDate

	nationality, err := formatter.field(nationalityField, secondLine, 15, 3, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.Nationality] = nationality

	optionalData2, err := formatter.field(optionalDataField, secondLine, 18, 11, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.OptionalData2] = optionalData2

	finalCheckDigit, err := formatter.field(hashField, secondLine, 29, 1, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.FinalCheckDigit] = finalCheckDigit

	// Parse third line
	name, err := formatter.field(namesField, thirdLine, 0, 30, false)
	if err != nil {
		return result, err
	}
	parsedResult[constants.Name] = name

	isValid, err := td1.validateAllCheckDigits(documentNumber, optionalData1, birthdate, expiryDate, optionalData2, finalCheckDigit)
	if err != nil {
		return result, err
	}

	fmt.Println(documentType, countryCode, documentNumber, optionalData1)
	fmt.Println(birthdate, sex, expiryDate, nationality, optionalData2, finalCheckDigit)
	fmt.Println(name)
	fmt.Println(isValid)

	result.Mapper = parsedResult
	result.IsValid = isValid

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
	if calculatedCheckDigit != finalCheckDigit.value {
		return false, nil
	}
	return true, nil
}
