package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
	"strings"
)

func main() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values := utils.SpliceByLine(file)
	lineSegments := readLines(values)

	fmt.Printf("Part1- %v\n", part1(lineSegments))
	fmt.Printf("Part2- %v\n", part2(lineSegments))
}

func part1(lines []lineSegment) int {
	board := [1000][1000]int{}

	for _, line := range lines {
		points := pointsBetween(line, true)
		for _, p := range points {
			board[p.x][p.y]++
		}
	}

	return countBoard(board)
}
func part2(lines []lineSegment) int {
	board := [1000][1000]int{}

	for _, line := range lines {
		points := pointsBetween(line, false)
		for _, p := range points {
			board[p.x][p.y]++
		}
	}

	return countBoard(board)
}

func pointsBetween(s lineSegment, onlyStraight bool) []point {

	var xs []int
	if s.start.x <= s.end.x {
		for x := s.start.x; x <= s.end.x; x++ {
			xs = append(xs, x)
		}
	} else {
		for x := s.start.x; x >= s.end.x; x-- {
			xs = append(xs, x)
		}
	}
	var ys []int
	if s.start.y <= s.end.y {
		for y := s.start.y; y <= s.end.y; y++ {
			ys = append(ys, y)
		}
	} else {
		for y := s.start.y; y >= s.end.y; y-- {
			ys = append(ys, y)
		}
	}

	var points []point

	if len(xs) == 1 {
		for _, y := range ys {
			points = append(points, point{
				x: xs[0],
				y: y,
			})
		}
	} else if len(ys) == 1 {
		for _, x := range xs {
			points = append(points, point{
				x: x,
				y: ys[0],
			})
		}
	} else if !onlyStraight {
		for i := 0; i < len(ys); i++ {
			points = append(points, point{
				x: xs[i],
				y: ys[i],
			})
		}
	}

	return points
}

func countBoard(board [1000][1000]int) int {
	var count int
	for _, i := range board {
		for _, value := range i {
			if value > 1 {
				count++
			}
		}
	}
	return count
}

type point struct {
	x int
	y int
}

type lineSegment struct {
	start point
	end   point
}

func readLines(lines []string) []lineSegment {
	var lineSegments []lineSegment
	for _, value := range lines {
		coordinates := strings.Split(value, " -> ")
		coordinates1, err := utils.SpliceByToInt(coordinates[0], ",")
		if err != nil {
			panic("")
		}
		startPoint := point{x: coordinates1[0], y: coordinates1[1]}
		coordinates2, err := utils.SpliceByToInt(coordinates[1], ",")
		if err != nil {
			panic("")
		}
		endPoint := point{x: coordinates2[0], y: coordinates2[1]}

		lineSegments = append(lineSegments, lineSegment{
			start: startPoint,
			end:   endPoint,
		})
	}
	return lineSegments
}
