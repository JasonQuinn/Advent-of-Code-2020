package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
	"regexp"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values := utils.SpliceByLine(file)

	inputs := readLines(values)

	fmt.Printf("Part1- %v\n", part1(inputs))
	fmt.Printf("Part2-\n")
	part2(inputs)
}

type points struct {
	x int
	y int
}

type fold struct {
	direction string
	line      int
}

type input struct {
	values []points
	folds  []fold
}

func readLines(lines []string) input {
	valueRegex, _ := regexp.Compile("(\\d+),(\\d+)")
	foldRegex, _ := regexp.Compile("fold along ([xy])=(\\d+)")

	input := input{}
	for _, line := range lines {
		if valueRegex.MatchString(line) {
			matches := valueRegex.FindStringSubmatch(line)
			input.values = append(input.values, points{x: utils.StringToInt(matches[1]), y: utils.StringToInt(matches[2])})
		} else if foldRegex.MatchString(line) {
			matches := foldRegex.FindStringSubmatch(line)
			input.folds = append(input.folds, fold{direction: matches[1], line: utils.StringToInt(matches[2])})
		}
	}
	return input
}

func part1(input input) int {
	paper := createPaper(input)

	input.folds = input.folds[:1]

	newPaper := foldPaper(input, paper)

	sum := 0
	for y := range newPaper {
		for x := range newPaper[y] {
			if newPaper[y][x] == "■" {
				sum++
			}
		}
	}

	return sum
}

func part2(input input) {
	paper := createPaper(input)

	newPaper := foldPaper(input, paper)

	printPaper(newPaper)
}

func foldPaper(input input, paper [][]string) [][]string {
	newPaper := paper
	for _, f := range input.folds {
		if f.direction == "x" {
			newPaper = foldX(newPaper, f.line)
		} else if f.direction == "y" {
			newPaper = foldY(newPaper, f.line)
		}
	}
	return newPaper
}

func foldY(paper [][]string, line int) [][]string {
	newPaper := make([][]string, line)
	for y := 0; y < len(paper); y++ {
		for x := 0; x < len(paper[y]); x++ {
			if y < line {
				if newPaper[y] == nil {
					newPaper[y] = make([]string, len(paper[y]))
				}
				newPaper[y][x] = paper[y][x]
			}
			if y > line {
				if paper[y][x] == "■" {
					newPaper[line-(y-line)][x] = paper[y][x]
				}
			}
		}
	}

	return newPaper
}

func foldX(paper [][]string, line int) [][]string {
	newPaper := make([][]string, len(paper))
	for y := 0; y < len(paper); y++ {
		for x := 0; x < len(paper[y]); x++ {
			if x < line {
				if newPaper[y] == nil {
					newPaper[y] = make([]string, line)
				}
				newPaper[y][x] = paper[y][x]
			}
			if x > line {
				if paper[y][x] == "■" {
					newPaper[y][line-(x-line)] = paper[y][x]
				}
			}
		}
	}

	return newPaper
}

func createPaper(input input) [][]string {
	var maxX int
	var maxY int
	for _, value := range input.values {
		if maxX < value.x {
			maxX = value.x
		}
		if maxY < value.y {
			maxY = value.y
		}
	}
	paper := make([][]string, maxY+1)
	for y := range paper {
		paper[y] = make([]string, maxX+1)
		for x := range paper[y] {
			paper[y][x] = " "
		}
	}

	for _, value := range input.values {
		paper[value.y][value.x] = "■"
	}

	return paper
}

func printPaper(paper [][]string) {
	for y := range paper {
		for x := range paper[y] {
			fmt.Print(paper[y][x])
		}
		fmt.Printf("\n")
	}
}
