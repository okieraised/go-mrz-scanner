package parser

// MRZResult defines the structure of the parsed MRZ string.
//   - IsVISA (bool): Determines if the MRZ string is VISA document.
//   - IsValid (bool): Determines if the MRZ string content is valid.
//   - Fields (map): Map of the MRZ fields.
type MRZResult struct {
	IsVISA  bool
	IsValid bool
	Fields  map[string]*mrzField
}

type IMRZParser interface {
	Parse(in []string) (*MRZResult, error)
}
