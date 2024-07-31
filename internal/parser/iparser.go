package parser

type ParserResult struct {
	IsVISA  bool
	IsValid bool
	Mapper  map[string]*mrzField
}

type IMRZParser interface {
	Parse(in []string) (*ParserResult, error)
}
