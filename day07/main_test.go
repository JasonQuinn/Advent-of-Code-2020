package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 37, part1([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 168, part2([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}))
}
