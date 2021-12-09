package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadLines(t *testing.T) {
	lines := readLines([]string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	})

	assert.Equal(t, 5, len(lines))
	assert.Equal(t, 10, len(lines[0]))
}

func TestPart1(t *testing.T) {
	lines := readLines([]string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	})

	assert.Equal(t, 15, part1(lines))
}

func TestPart2(t *testing.T) {
	lines := readLines([]string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	})

	assert.Equal(t, 1134, part2(lines))
}
