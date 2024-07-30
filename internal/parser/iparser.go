package parser

type ParserResult struct {
	IsValid bool
	Mapper  map[string]*mrzField
}

type IMRZParser interface {
	Parse(in []string) (*ParserResult, error)
}
