package go_mrz_parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMRZParser_Parse_TD1_String(t *testing.T) {
	parser := NewMRZStringParser(
		"I<UTOD231458907<<<<<<<<<<<<<<<7408122F1204159UTO<<<<<<<<<<<6ERIKSSON<<ANNA<MARIA<<<<<<<<<<",
	)

	result, err := parser.Parse()
	assert.NoError(t, err)
	fmt.Println(result)
}

func TestMRZParser_Parse_TD1_StringSlice(t *testing.T) {
	parser := NewMRZLineParser([]string{
		"I<UTOD231458907<<<<<<<<<<<<<<<",
		"7408122F1204159UTO<<<<<<<<<<<6",
		"ERIKSSON<<ANNA<MARIA<<<<<<<<<<",
	})

	result, err := parser.Parse()
	assert.NoError(t, err)
	fmt.Println(result)
}

func TestMRZParser_Parse_TD2_String(t *testing.T) {
	parser := NewMRZStringParser(
		"I<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<D231458907UTO7408122F1204159<<<<<<<6",
	)

	result, err := parser.Parse()
	assert.NoError(t, err)
	fmt.Println(result)
}

func TestMRZParser_Parse_TD2_StringSlice(t *testing.T) {
	parser := NewMRZLineParser([]string{
		"I<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<",
		"D231458907UTO7408122F1204159<<<<<<<6",
	})

	result, err := parser.Parse()
	assert.NoError(t, err)
	fmt.Println(result)
}

func TestMRZParser_Parse_TD3_String(t *testing.T) {
	parser := NewMRZStringParser(
		"P<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<<<<<<<<<L898902C36UTO7408122F1204159ZE184226B<<<<<10",
	)

	result, err := parser.Parse()
	assert.NoError(t, err)
	assert.Equal(t, false, result.IsVISA)
	assert.Equal(t, true, result.IsValid)
}

func TestMRZParser_Parse_TD3_StringSlice(t *testing.T) {
	parser := NewMRZLineParser([]string{
		"P<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<<<<<<<<<",
		"L898902C36UTO7408122F1204159ZE184226B<<<<<10",
	})

	result, err := parser.Parse()
	assert.NoError(t, err)
	assert.Equal(t, false, result.IsVISA)
	assert.Equal(t, true, result.IsValid)
}

func TestMRZParser_Parse_TD3_VISA(t *testing.T) {
	parser := NewMRZLineParser([]string{
		"V<UTOERIKSSON<<ANNA<MARIA<<<<<<<<<<<<<<<<<<<",
		"L8988901C4XXX4009078F96121096ZE184226B<<<<<<",
	})

	result, err := parser.Parse()
	assert.NoError(t, err)
	assert.Equal(t, true, result.IsVISA)
	assert.Equal(t, true, result.IsValid)
}
