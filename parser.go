package go_mrz_parser

import (
	"fmt"
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

func NewMRZLineParser(mrzLines []string) *MRZParser {
	return &MRZParser{
		components: mrzLines,
	}
}

func (p *MRZParser) Parse() error {

	err := p.validate()
	if err != nil {
		return err
	}

	var mrzParser parser.IMRZParser
	switch p.mrzType {
	case constants.MRZType1:
		mrzParser = parser.NewTD1()
	case constants.MRZType2:
		mrzParser = parser.NewTD2()
	case constants.MRZType3:
		mrzParser = parser.NewTD3()
	}

	parse, err := mrzParser.Parse(p.components)
	if err != nil {
		return err
	}

	fmt.Println(parse)

	return nil
}

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
		characterCount := make([]int, 2)
		for _, line := range p.components {
			if len(line) != constants.Type2NumberOfCharacter &&
				len(line) != constants.Type3NumberOfCharacter {
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
