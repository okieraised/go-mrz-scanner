package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateCheckDigits(t *testing.T) {
	character := "1234567889abcdef"

	checkDigit, err := CalculateCheckDigits(character)
	assert.Nil(t, err)
	assert.Equal(t, "2", checkDigit)
}
