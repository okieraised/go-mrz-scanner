package go_mrz_parser

import (
	"github.com/okieraised/go-mrz-scanner/constants"
	"github.com/okieraised/go-mrz-scanner/internal/parser"
	"github.com/okieraised/go-mrz-scanner/mrz_errors"
	"github.com/okieraised/go-mrz-scanner/utils"
	"strings"
)

type MRZParser struct {
	mrzType    int
	components []string
}

// NewMRZStringParser receives a single mrz string with each line separated by newline character
func NewMRZStringParser(mrzStr string) *MRZParser {
	components := strings.Split(mrzStr, "\n")
	return &MRZParser{
		components: components,
	}
}

// NewMRZLineParser receives the mrz string slices
func NewMRZLineParser(mrzLines []string) *MRZParser {
	return &MRZParser{
		components: mrzLines,
	}
}

// Parse validates and parses the MRZ information
func (p *MRZParser) Parse() (*parser.ParserResult, error) {

	err := p.validate()
	if err != nil {
		return &parser.ParserResult{}, err
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
		return &parser.ParserResult{}, mrz_errors.ErrInvalidMRZType
	}

	parse, err := mrzParser.Parse(p.components)
	if err != nil {
		return &parser.ParserResult{}, err
	}

	return parse, nil
}

// validate checks the input MRZ for formatting errors
func (p *MRZParser) validate() error {
	mrzType := 0

	switch len(p.components) {
	case 3:
		for _, line := range p.components {
			if len(line) != constants.Type1NumberOfCharacter {
				return mrz_errors.ErrTD1InvalidLineLength
			}
		}
		mrzType = constants.MRZType1
	case 2:
		characterCount := make([]int, 0)
		for _, line := range p.components {
			if len(line) != constants.Type2NumberOfCharacter && len(line) != constants.Type3NumberOfCharacter {
				return mrz_errors.ErrGenericInvalidMRZLinesLength
			}
			characterCount = append(characterCount, len(line))
		}
		if utils.CheckSame(characterCount) && characterCount[0] == constants.Type2NumberOfCharacter {
			mrzType = constants.MRZType2
		}
		if utils.CheckSame(characterCount) && characterCount[0] == constants.Type3NumberOfCharacter {
			mrzType = constants.MRZType3
		}
	default:
		return mrz_errors.ErrGenericInvalidMRZLines
	}
	p.mrzType = mrzType
	return nil
}
