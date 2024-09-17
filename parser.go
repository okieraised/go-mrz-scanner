package go_mrz_parser

import (
	"github.com/okieraised/go-mrz-scanner/constants"
	"github.com/okieraised/go-mrz-scanner/internal/parser"
	"github.com/okieraised/go-mrz-scanner/mrz_errors"
	"github.com/okieraised/go-mrz-scanner/utils"
	"strings"
)

// MRZParser defines the structure of the parser.
//   - mrzType: type of the MRZ string
//   - components: parts of the MRZ string
type MRZParser struct {
	mrzType    int
	components []string
}

// NewMRZStringParser receives a single mrz string with each line separated by newline character.
func NewMRZStringParser(mrzStr string) *MRZParser {
	components := make([]string, 0)
	if strings.Contains(mrzStr, "\n") {
		components = strings.Split(mrzStr, "\n")
	} else {
		if len(mrzStr) == constants.Type1TotalNumberOfCharacters {
			components = []string{
				mrzStr[:constants.Type1NumberOfCharactersPerLine],
				mrzStr[constants.Type1NumberOfCharactersPerLine : 2*constants.Type1NumberOfCharactersPerLine],
				mrzStr[2*constants.Type1NumberOfCharactersPerLine:],
			}
		} else if len(mrzStr) == constants.Type2TotalNumberOfCharacters {
			components = []string{
				mrzStr[:constants.Type2NumberOfCharactersPerLine],
				mrzStr[constants.Type2NumberOfCharactersPerLine:],
			}
		} else {
			components = []string{
				mrzStr[:constants.Type3NumberOfCharactersPerLine],
				mrzStr[constants.Type3NumberOfCharactersPerLine:],
			}
		}
	}

	return &MRZParser{
		components: components,
	}
}

// NewMRZLineParser receives the mrz string slices.
func NewMRZLineParser(mrzLines []string) *MRZParser {
	return &MRZParser{
		components: mrzLines,
	}
}

// Parse validates and parses the MRZ information
func (p *MRZParser) Parse() (*parser.MRZResult, error) {

	err := p.validate()
	if err != nil {
		return &parser.MRZResult{}, err
	}

	var mrzParser parser.IMRZParser
	switch p.mrzType {
	case constants.MRZType1:
		mrzParser = parser.NewTD1()
	case constants.MRZType2:
		mrzParser = parser.NewTD2()
	case constants.MRZType3:
		mrzParser = parser.NewTD3()
	default:
		return &parser.MRZResult{}, mrz_errors.ErrInvalidMRZType
	}

	parse, err := mrzParser.Parse(p.components)
	if err != nil {
		return &parser.MRZResult{}, err
	}

	return parse, nil
}

// validate checks the input MRZ for formatting errors
func (p *MRZParser) validate() error {
	mrzType := 0

	switch len(p.components) {
	case 3:
		for _, line := range p.components {
			if len(line) != constants.Type1NumberOfCharactersPerLine {
				return mrz_errors.ErrTD1InvalidLineLength
			}
		}
		mrzType = constants.MRZType1
	case 2:
		characterCount := make([]int, 0)
		for _, line := range p.components {
			if len(line) != constants.Type2NumberOfCharactersPerLine && len(line) != constants.Type3NumberOfCharactersPerLine {
				return mrz_errors.ErrGenericInvalidMRZLinesLength
			}
			characterCount = append(characterCount, len(line))
		}
		if utils.CheckSame(characterCount) && characterCount[0] == constants.Type2NumberOfCharactersPerLine {
			mrzType = constants.MRZType2
		}
		if utils.CheckSame(characterCount) && characterCount[0] == constants.Type3NumberOfCharactersPerLine {
			mrzType = constants.MRZType3
		}
	default:
		return mrz_errors.ErrGenericInvalidMRZLines
	}
	p.mrzType = mrzType
	return nil
}
