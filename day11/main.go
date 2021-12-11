package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values := utils.SpliceByLine(file)

	fmt.Printf("Part1- %v\n", part1(readLines(values), 100))
	fmt.Printf("Part2- %v\n", part2(readLines(values)))
}

func part1(board [][]int, numberOfTimes int) int {
	defer utils.Duration(utils.Track("part1"))
	//Increase each
	numberFires := 0
	for i := 1; i < numberOfTimes; i++ {
		for x, row := range board {
			for y := range row {
				numberFires += increaseOctopus(board, x, y)
			}
		}

		for x, row := range board {
			for y := range row {
				if board[x][y] > 9 {
					board[x][y] = 0
				}
			}
		}
	}
	return numberFires
}

func part2(board [][]int) int {
	defer utils.Duration(utils.Track("part1"))
	//Increase each
	for i := 1; true; i++ {
		for x, row := range board {
			for y := range row {
				increaseOctopus(board, x, y)
			}
		}

		allFlash := true
		for x, row := range board {
			for y := range row {
				if board[x][y] > 9 {
					board[x][y] = 0
				} else {
					allFlash = false
				}
			}
		}
		if allFlash {
			return i
		}
	}
	return -1
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

func increaseOctopus(board [][]int, x int, y int) int {
	numberFired := 0
	if x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) {
		if board[x][y] <= 10 {
			board[x][y]++
		}
		if board[x][y] == 10 {
			numberFired++
			//fire each surrounding
			numberFired += increaseOctopus(board, x-1, y)
			numberFired += increaseOctopus(board, x-1, y-1)
			numberFired += increaseOctopus(board, x, y-1)
			numberFired += increaseOctopus(board, x+1, y)
			numberFired += increaseOctopus(board, x+1, y+1)
			numberFired += increaseOctopus(board, x, y+1)
			numberFired += increaseOctopus(board, x+1, y-1)
			numberFired += increaseOctopus(board, x-1, y+1)
		}
	}
	return numberFired
}
