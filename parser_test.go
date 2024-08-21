package go_mrz_parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMRZParser_Parse_TD1(t *testing.T) {
	parser := NewMRZLineParser([]string{
		"I<UTOD231458907<<<<<<<<<<<<<<<",
		"7408122F1204159UTO<<<<<<<<<<<6",
		"ERIKSSON<<ANNA<MARIA<<<<<<<<<<",
	})

	err := parser.Parse()
	assert.NoError(t, err)
}

func TestMRZParser_Parse_TD2(t *testing.T) {
	parser := NewMRZLineParser([]string{
		"I<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<",
		"D231458907UTO7408122F1204159<<<<<<<6",
	})

	err := parser.Parse()
	assert.NoError(t, err)
}
