package mrz_errors

import "errors"

var (
	ErrInvalidMRZType = errors.New("invalid mrz type")
)

var (
	ErrGenericInvalidMRZLength      = errors.New("invalid mrz length")
	ErrGenericInvalidMRZLines       = errors.New("invalid number of mrz lines")
	ErrGenericInvalidMRZLinesLength = errors.New("invalid mrz line length")
	ErrTD1InvalidLineLength         = errors.New("invalid TD1 format line length")
	ErrTD2InvalidLineLength         = errors.New("invalid TD2 format line length")
	ErrTD3InvalidLineLength         = errors.New("invalid TD3 format line length")
)

var (
	ErrInvalidBirthdateCharacter = errors.New("invalid character in birthdate field")
)
