package constants

const (
	MRZType1 = iota + 1
	MRZType2
	MRZType3
)

const (
	Type1NumberOfCharacter = 30
	Type2NumberOfCharacter = 36
	Type3NumberOfCharacter = 44
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
