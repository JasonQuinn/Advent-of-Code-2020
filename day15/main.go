package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"math"
	"os"
)

var min int

func main() {
	min = math.MaxInt

	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values := utils.SpliceByLine(file)

	board := readLines(values)

	fmt.Printf("Part1- %v\n", part1(board))

	fmt.Printf("Part1- %v\n", part1(getLargerBoard(board)))
}

func getLargerBoard(board [][]int) [][]int {
	newBoard := make([][]int, len(board)*5)

	for y := range newBoard {
		newRow := make([]int, len(board[0])*5)
		newBoard[y] = newRow
	}

	for boardY := 0; boardY < len(board); boardY++ {
		for boardX := 0; boardX < len(board[boardY]); boardX++ {
			for yMultiplies := 0; yMultiplies < 5; yMultiplies++ {
				for xMultiplies := 0; xMultiplies < 5; xMultiplies++ {

					newValue := board[boardY][boardX] + yMultiplies + xMultiplies
					if newValue >= 10 {
						newValue = newValue%10 + 1
					}

					newBoard[boardY+(len(board)*yMultiplies)][boardX+(len(board[boardY])*xMultiplies)] = newValue
				}
			}
		}
	}

	return newBoard
}

type cell struct {
	x, y int
}

func part1(matrix [][]int) int {

	distances := make([][]int, len(matrix))
	for i := range distances {
		aux := make([]int, len(matrix[0]))
		for j := range aux {
			aux[j] = math.MaxInt
		}
		distances[i] = aux
	}

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	setCells := make(map[cell]int)
	setCells[cell{x: 0, y: 0}] = 0

	distances[0][0] = matrix[0][0]

	for len(setCells) != 0 {

		var cell0 cell
		for k := range setCells {
			cell0 = k
			break
		}
		delete(setCells, cell0)

		for i := 0; i < 4; i++ {
			x := cell0.x + dx[i]
			y := cell0.y + dy[i]

			if !isInsideGrid(matrix, x, y) {
				continue
			} else if distances[x][y] > distances[cell0.x][cell0.y]+matrix[x][y] {
				if distances[x][y] != math.MaxInt {
					if _, ok := setCells[cell0]; ok {
						delete(setCells, cell0)
						setCells[cell0] = distances[cell0.x][cell0.y] + matrix[x][y]
					}
				}
				distances[x][y] = distances[cell0.x][cell0.y] + matrix[x][y]
				setCells[cell{x: x, y: y}] = distances[x][y]
			}
		}

	}

	return distances[len(matrix)-1][len(matrix[0])-1] - matrix[0][0]

}

func isInsideGrid(matrix [][]int, i, j int) bool {
	return i >= 0 && i < int(len(matrix)) && j >= 0 && j < int(len(matrix[0]))
}

func Min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
func readLines(lines []string) [][]int {
	var board [][]int
	for _, line := range lines {
		var boardLine []int
		for _, c := range line {
			boardLine = append(boardLine, utils.StringToInt(string(c)))
		}
		board = append(board, boardLine)
	}
	return board
}
