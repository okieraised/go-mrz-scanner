package constants

const (
	MRZType1 = iota + 1
	MRZType2
	MRZType3
)

const (
	Type1TotalNumberOfCharacters = 3 * Type1NumberOfCharactersPerLine
	Type2TotalNumberOfCharacters = 2 * Type2NumberOfCharactersPerLine
	Type3TotalNumberOfCharacters = 2 * Type3NumberOfCharactersPerLine
)

const (
	Type1NumberOfCharactersPerLine = 30
	Type2NumberOfCharactersPerLine = 36
	Type3NumberOfCharactersPerLine = 44
)

const (
	MRZFieldDocumentType    = "document_type"
	MRZFieldCountryCode     = "country_code"
	MRZFieldDocumentNumber  = "document_number"
	MRZFieldOptionalData1   = "optional_data_1"
	MRZFieldBirthdate       = "birthdate"
	MRZFieldGender          = "sex"
	MRZFieldExpiryDate      = "expiry_date"
	MRZFieldNationality     = "nationality"
	MRZFieldOptionalData2   = "optional_data_2"
	MRZFieldFinalCheckDigit = "final_check_digit"
	MRZFieldName            = "name"
	MRZFieldSurname         = "surname"
	MRZFieldGivenName       = "given_name"
)
