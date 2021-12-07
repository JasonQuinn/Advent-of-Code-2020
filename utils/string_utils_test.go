package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// BinaryToInt
func TestBinaryStringToInt(t *testing.T) {
	assert.Equal(t, 10, BinaryStringToInt("01010"))
	assert.Equal(t, 23, BinaryStringToInt("10111"))
	assert.Equal(t, 9, BinaryStringToInt("01001"))
	assert.Equal(t, 22, BinaryStringToInt("10110"))
}
