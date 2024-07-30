package go_mrz_parser

// WithSingleMRZ received the mrz as a single string
func WithSingleMRZ(mrzString string) func(*MRZParser) {
	return func(c *MRZParser) {
		c.singleString = mrzString
	}
}

// WithMRZComponents received the mrz string as individual components
func WithMRZComponents(components []string) func(*MRZParser) {
	return func(c *MRZParser) {
		c.componentStrings = components
	}
}

type MRZParser struct {
	singleString     string
	componentStrings []string
}

func NewParser(options ...func(*MRZParser)) *MRZParser {
	mrzParser := &MRZParser{}

	for _, opt := range options {
		opt(mrzParser)
	}
	return mrzParser
}

func (p *MRZParser) Parse() {

}
