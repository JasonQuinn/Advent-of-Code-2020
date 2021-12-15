package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadLine(t *testing.T) {
	board := readLines([]string{
		"1163751742",
		"1381373672",
		"2136511328",
		"3694931569",
		"7463417111",
		"1319128137",
		"1359912421",
		"3125421639",
		"1293138521",
		"2311944581",
	})

	assert.Equal(t, 10, len(board))
}

func TestPart1(t *testing.T) {
	board := readLines([]string{
		"1163751742",
		"1381373672",
		"2136511328",
		"3694931569",
		"7463417111",
		"1319128137",
		"1359912421",
		"3125421639",
		"1293138521",
		"2311944581",
	})

	assert.Equal(t, 40, part1(board))
}

func TestLargerBoard(t *testing.T) {
	board := readLines([]string{
		"1163751742",
		"1381373672",
		"2136511328",
		"3694931569",
		"7463417111",
		"1319128137",
		"1359912421",
		"3125421639",
		"1293138521",
		"2311944581",
	})

	newBoard := getLargerBoard(board)
	assert.Equal(t, 50, len(newBoard))
}

func TestPart2(t *testing.T) {
	board := readLines([]string{
		"1163751742",
		"1381373672",
		"2136511328",
		"3694931569",
		"7463417111",
		"1319128137",
		"1359912421",
		"3125421639",
		"1293138521",
		"2311944581",
	})

	assert.Equal(t, 315, part1(getLargerBoard(board)))
}
