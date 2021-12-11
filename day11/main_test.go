package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	board := readLines([]string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	})
	//assert.Equal(t, 0, part1(board, 1))
	assert.Equal(t, 204, part1(board, 10))
}

func TestPart2(t *testing.T) {
	board := readLines([]string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	})
	assert.Equal(t, 195, part2(board))
}
