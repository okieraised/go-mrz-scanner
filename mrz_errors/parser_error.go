package mrz_errors

import "errors"

var (
	ErrGenericInvalidMRZLength = errors.New("invalid mrz length")
	ErrTD1InvalidLineLength    = errors.New("invalid TD1 format line length")
)

var (
	ErrInvalidBirthdateCharacter = errors.New("invalid character in birthdate field")
)
