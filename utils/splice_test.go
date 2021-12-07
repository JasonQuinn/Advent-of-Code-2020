package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInputToSpliceEmpty(t *testing.T) {
	values, err := SpliceByLine("")
	assert.Nil(t, err)
	assert.Len(t, values, 1)
}

func TestInputToSplice(t *testing.T) {
	values, err := SpliceByLine("1\n2")
	assert.Nil(t, err)
	assert.Len(t, values, 2)

	assert.ElementsMatch(t, []string{"1", "2"}, values)
}
