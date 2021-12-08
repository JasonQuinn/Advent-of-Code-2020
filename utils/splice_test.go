package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInputToSpliceEmpty(t *testing.T) {
	assert.Len(t, SpliceByLine(" "), 0)
	assert.Len(t, SpliceByLine(""), 0)
}

func TestInputToSplice(t *testing.T) {
	values := SpliceByLine("1\n2")
	assert.Len(t, values, 2)

	assert.ElementsMatch(t, []string{"1", "2"}, values)
}

func TestSpliceByToInt(t *testing.T) {
	values, err := SpliceByToInt("1,2,4", ",")
	assert.Nil(t, err)
	assert.Len(t, values, 3)
	assert.ElementsMatch(t, []int{1, 2, 4}, values)
}

func TestSpliceBy(t *testing.T) {
	values := SpliceBy("1 2", " ")
	assert.Len(t, values, 2)

	assert.ElementsMatch(t, []string{"1", "2"}, values)
}
