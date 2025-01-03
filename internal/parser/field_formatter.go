package parser

import (
	"github.com/okieraised/go-mrz-scanner/constants"
	"github.com/okieraised/go-mrz-scanner/mrz_errors"
	"github.com/okieraised/go-mrz-scanner/utils"
	"strconv"
	"strings"
	"time"
)

const (
	documentTypeField = iota
	countryCodeField
	namesField
	documentNumberField
	nationalityField
	birthdateField
	sexField
	expiryDateField
	personalNumberField
	optionalDataField
	alphabeticField
	hashField
	numericField
)

const (
	mrzDateFmt = "060102" // YYMMDD
)

// mrzField specifies the structure for each field in the MRZ string.
//
//   - value: Parsed value (useful for name field)
//   - rawValue: Raw MRZ value as string
//   - checkDigit: Check digit (if available)
//   - isValid: Specifies if the field value is valid
type mrzField struct {
	value      any
	rawValue   string
	checkDigit string
	isValid    bool
}

// GetNames returns the name field as a string slice.
func (f *mrzField) GetNames() ([]string, error) {
	names, ok := f.value.([]string)
	if !ok {
		return nil, mrz_errors.ErrNotNameField
	}
	return names, nil
}

// GetValue returns the parsed value.
func (f *mrzField) GetValue() any {
	return f.value
}

// GetRawValue returns the field value as string.
func (f *mrzField) GetRawValue() string {
	return f.rawValue
}

// GetCheckDigit returns the check digit value as string.
func (f *mrzField) GetCheckDigit() string {
	return f.checkDigit
}

// isValueValid performs field validity check.
func (f *mrzField) isValueValid() bool {
	if f.checkDigit == "<" {
		if len(utils.TrimmingFiller(f.rawValue)) != 0 {
			return false
		}
		f.checkDigit = "0"
	} else {
		_, err := strconv.Atoi(f.checkDigit)
		if err != nil {
			return false
		}
	}
	return utils.IsValueValid(f.rawValue, f.checkDigit)
}

type FieldFormatter struct {
	ocrCorrection bool
}

func NewFieldFormatter(ocrCorrection bool) *FieldFormatter {
	return &FieldFormatter{
		ocrCorrection: ocrCorrection,
	}
}

func (f *FieldFormatter) field(fieldType int, from string, startIdx, length int, checkDigitFollow bool) (*mrzField, error) {
	endIdx := startIdx + length
	rawValue := from[startIdx:endIdx]
	checkDigit := ""
	if checkDigitFollow {
		checkDigit = from[endIdx : endIdx+1]
	}

	if f.ocrCorrection {
		rawValue = f.correct(rawValue, fieldType)
	}

	fmtVal, err := f.format(rawValue, fieldType)
	if err != nil {
		return nil, err
	}

	field := &mrzField{
		value:      fmtVal,
		rawValue:   rawValue,
		checkDigit: checkDigit,
		isValid:    true,
	}

	if checkDigitFollow {
		field.isValid = field.isValueValid()
	}

	return field, nil

}

func (f *FieldFormatter) sex(from string) string {

	switch from {
	case "M":
		return "MALE"
	case "F":
		return "FEMALE"
	case "<":
		return "UNSPECIFIED" // X
	default:
		return ""
	}
}

func (f *FieldFormatter) names(from string) []string {

	var primary, secondary string
	identifiers := strings.Split(from, "<<")
	primary = strings.ReplaceAll(identifiers[0], "<", " ")
	if len(identifiers) == 1 {
		secondary = ""
	} else {
		secondary = strings.ReplaceAll(identifiers[1], "<", " ")
	}

	return []string{primary, secondary}

}

func (f *FieldFormatter) text(from string) string {
	return strings.ReplaceAll(from, "<", " ")
}

func (f *FieldFormatter) date(from string) (string, error) {

	for _, digit := range from {
		if string(digit) == "<" {
			continue
		}
		if !strings.Contains(constants.DecimalDigits, string(digit)) {
			return "", mrz_errors.ErrInvalidBirthdateCharacter
		}
	}

	if strings.Contains(from, "<") {
		return from, nil
	}
	_, err := time.Parse(mrzDateFmt, from)
	if err != nil {
		return "", err
	}

	return from, nil
}

func (f *FieldFormatter) replaceDigits(from string) string {
	return utils.ReplaceDigits(from)
}

func (f *FieldFormatter) replaceLetters(from string) string {
	return utils.ReplaceLetters(from)
}

func (f *FieldFormatter) correct(from string, fieldType int) string {
	switch fieldType {
	case birthdateField, expiryDateField, hashField, numericField:
		return f.replaceLetters(from)
	case namesField, documentTypeField, countryCodeField, nationalityField, alphabeticField:
		return f.replaceDigits(from)
	case sexField:
		return strings.ReplaceAll(from, "P", "F")
	default:
		return from
	}
}

func (f *FieldFormatter) format(from string, fieldType int) (any, error) {
	switch fieldType {
	case namesField:
		return f.names(from), nil
	case birthdateField, expiryDateField:
		return f.date(from)
	case sexField:
		return f.sex(from), nil
	default:
		return from, nil
	}
}
