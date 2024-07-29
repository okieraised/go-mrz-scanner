package parser

type TD1 struct {
	lineLength int
}

func NewTD1() IMRZParser {
	return &TD1{
		lineLength: 30,
	}
}

func (td1 *TD1) Parse(in []string) error {
	//TODO implement me
	panic("implement me")
}
