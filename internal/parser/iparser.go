package parser

type ParserResult struct {
	IsVISA  bool
	IsValid bool
	Fields  map[string]*mrzField
}

type IMRZParser interface {
	Parse(in []string) (*ParserResult, error)
}
