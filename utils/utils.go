package utils

import (
	"errors"
	"github.com/okieraised/go-mrz-scanner/constants"
	"strconv"
	"strings"
)

func IsValueValid(value, checkDigit string) bool {
	var err error
	total := 0
	charValue := 0

	for idx, v := range strings.ToUpper(value) {
		if strings.Contains(constants.UpperCaseLetter, string(v)) {
			charValue = constants.AlphabetMapping[string(v)]
		} else if strings.Contains(constants.DecimalDigits, string(v)) {
			charValue, err = strconv.Atoi(string(v))
			if err != nil {
				return false
			}
		} else if string(v) == "<" {
			charValue = 0
		} else {
			return false
		}

		total += charValue * constants.Weights[idx%3]
	}
	if strconv.Itoa(total%10) != checkDigit {
		return false
	} else {
		return true
	}
}

func CalculateCheckDigits(value string) (string, error) {

	var err error
	total := 0
	charValue := 0

	for idx, v := range strings.ToUpper(value) {
		if strings.Contains(constants.UpperCaseLetter, string(v)) {
			charValue = constants.AlphabetMapping[string(v)]
		} else if strings.Contains(constants.DecimalDigits, string(v)) {
			charValue, err = strconv.Atoi(string(v))
			if err != nil {
				return "", err
			}
		} else if string(v) == "<" {
			charValue = 0
		} else {
			return "", errors.New("invalid mrz character")
		}

		total += charValue * constants.Weights[idx%3]
	}

	return strconv.Itoa(total % 10), nil
}

func ReplaceDigits(in string) string {
	replacer := strings.NewReplacer(
		"0", "O",
		"1", "I",
		"2", "Z",
		"8", "B",
	)

	return replacer.Replace(in)
}

func ReplaceLetters(in string) string {
	replacer := strings.NewReplacer(
		"O", "0",
		"Q", "0",
		"U", "0",
		"D", "0",
		"I", "1",
		"Z", "2",
		"B", "8",
	)

	return replacer.Replace(in)
}

func TrimmingFiller(in string) string {
	return strings.Trim(in, "<")
}
