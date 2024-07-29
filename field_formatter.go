package parser

const (
	documentType = iota
	countryCode
	names
	documentNumber
	nationality
	birthdate
	sex
	expiryDate
	personalNumber
	optionalData
	hash
)

type FieldFormatter struct {
	dateFormatter string
	ocrCorrection bool
}

func NewFieldFormatter() *FieldFormatter {
	return &FieldFormatter{}
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
