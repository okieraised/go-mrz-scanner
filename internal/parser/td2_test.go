package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTD2_Parse(t *testing.T) {
	parser := NewTD2()

	mapper, err := parser.Parse([]string{
		"I<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<",
		"D231458907UTO7408122F1204159<<<<<<<6",
	})

	assert.NoError(t, err)
	assert.Equal(t, false, mapper.IsVISA)
	assert.Equal(t, true, mapper.IsValid)
}

func TestTD2_Parse_VISA(t *testing.T) {
	parser := NewTD2()

	mapper, err := parser.Parse([]string{
		"V<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<",
		"L8988901C4XXX4009078F9612109<<<<<<<<",
	})

	assert.NoError(t, err)
	assert.Equal(t, true, mapper.IsVISA)
	assert.Equal(t, true, mapper.IsValid)
}
