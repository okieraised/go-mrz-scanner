package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTD1_Parse(t *testing.T) {
	parser := NewTD1()

	mapper, err := parser.Parse([]string{
		"I<UTOD231458907<<<<<<<<<<<<<<<",
		"7408122F1204159UTO<<<<<<<<<<<6",
		"ERIKSSON<<ANNA<MARIA<<<<<<<<<<",
	})

	assert.NoError(t, err)
	fmt.Println(mapper)
}
