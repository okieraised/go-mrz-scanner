package go_mrz_parser

type MRZInfo struct {
	documentType           string
	countryCode            string
	surnames               string
	givenNames             string
	documentNumber         string
	nationalityCountryCode string
	birthdate              string
	sex                    string
	expiryDate             string
	personalNumber         string
	personalNumber2        string
	isDocumentNumberValid  bool
	isBirthdateValid       bool
	isExpiryDateValid      bool
	isPersonalNumberValid  bool
	allCheckDigitsValid    bool
}

type MRZFields struct{}
