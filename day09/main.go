package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values := utils.SpliceByLine(file)
	lines := readLines(values)

	fmt.Printf("Part1- %v\n", part1(lines))
	fmt.Printf("Part2- %v\n", part2(lines))
}

func readLines(lines []string) [][]int {
	var linesArray [][]int
	for _, line := range lines {
		var lineArray []int
		for _, l := range line {
			number, err := strconv.Atoi(string(l))
			if err != nil {
				panic("Not a number")
			}
			lineArray = append(lineArray, number)
		}
		linesArray = append(linesArray, lineArray)
	}
	return linesArray
}

func part1(board [][]int) int {
	var lowPoints []int
	for x, column := range board {
		for y, value := range column {
			if (x-1 < 0 || value < board[x-1][y]) && (x+1 > len(board)-1 || value < board[x+1][y]) && (y-1 < 0 || value < board[x][y-1]) && (y+1 > len(column)-1 || value < board[x][y+1]) {
				lowPoints = append(lowPoints, value)
			}
		}
	}

	var risk int
	for _, lowPoint := range lowPoints {
		risk += lowPoint + 1
	}

	return risk
}

func part2(board [][]int) int {
	var sizes []int
	for x, row := range board {
		for y, value := range row {
			if (x-1 < 0 || value < board[x-1][y]) && (x+1 > len(board)-1 || value < board[x+1][y]) && (y-1 < 0 || value < board[x][y-1]) && (y+1 > len(row)-1 || value < board[x][y+1]) {
				sizes = append(sizes, workOutBasinSize(board, point{
					x: x,
					y: y,
				}))
			}
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes[0] * sizes[1] * sizes[2]
}

func mergeMap(map1 map[point]bool, map2 map[point]bool) map[point]bool {
	newMap := make(map[point]bool)
	for k, v := range map1 {
		newMap[k] = v
	}
	for k, v := range map2 {
		newMap[k] = v
	}
	return newMap
}

type point struct {
	x int
	y int
}

func workOutBasinSize(board [][]int, lowPoint point) int {
	pointsInRow := getToSides(board, lowPoint)

	below := getNextRow(board, pointsInRow, 1)

	above := getNextRow(board, pointsInRow, -1)

	newMap := mergeMap(pointsInRow, mergeMap(below, above))

	return len(newMap)
}

func getNextRow(board [][]int, currentPositions map[point]bool, direction int) map[point]bool {
	row := make(map[point]bool)
	for position := range currentPositions {
		valuesToSide := getToSides(board, point{
			x: position.x + direction,
			y: position.y,
		})
		for k, v := range valuesToSide {
			row[k] = v
		}
	}
	if len(row) == 0 {
		return row
	} else {
		return mergeMap(row, getNextRow(board, row, direction))
	}
}

func getToSides(board [][]int, currentPoint point) map[point]bool {
	if currentPoint.x < 0 || currentPoint.y < 0 || currentPoint.x >= len(board) || currentPoint.y >= len(board[currentPoint.x]) || board[currentPoint.x][currentPoint.y] == 9 {
		return nil
	}
	pointsInRow := make(map[point]bool)
	for y := currentPoint.y + 1; y < len(board[currentPoint.x]); y++ {
		if board[currentPoint.x][y] < 9 {
			pointsInRow[point{x: currentPoint.x, y: y}] = true
		} else {
			break
		}
	}
	if board[currentPoint.x][currentPoint.y] < 9 {
		pointsInRow[currentPoint] = true
	}
	for y := currentPoint.y - 1; y >= 0; y-- {
		if board[currentPoint.x][y] < 9 {
			pointsInRow[point{x: currentPoint.x, y: y}] = true
		} else {
			break
		}
	}
	return pointsInRow
}
