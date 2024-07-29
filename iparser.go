package parser

type IMRZParser interface {
	Parse(in []string) error
}
