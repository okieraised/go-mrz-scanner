package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTD3_Parse(t *testing.T) {
	parser := NewTD3()

	mapper, err := parser.Parse([]string{
		"P<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<<<<<<<<<",
		"L898902C36UTO7408122F1204159ZE184226B<<<<<10",
	})

	assert.NoError(t, err)
	assert.Equal(t, false, mapper.IsVISA)
	assert.Equal(t, true, mapper.IsValid)
}

func TestTD3_Parse_VISA(t *testing.T) {
	parser := NewTD3()

	mapper, err := parser.Parse([]string{
		"V<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<<<<<<<<<",
		"L8988901C4XXX4009078F96121096ZE184226B<<<<<<",
	})

	assert.NoError(t, err)
	//assert.Equal(t, true, mapper.IsVISA)
	//assert.Equal(t, true, mapper.IsValid)
	fmt.Println(mapper.IssuingState)
}
