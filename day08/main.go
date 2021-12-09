package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values := utils.SpliceByLine(file)
	lines := readLine(values)

	fmt.Printf("Part1- %v\n", part1(lines))
	fmt.Printf("Part2- %v\n", part2(lines))
}

type line struct {
	signal [10]string
	output [4]string
}

func readLine(input []string) []line {
	var lines []line
	for _, value := range input {
		splitLine := strings.Split(value, " | ")
		signal := utils.SpliceBy(splitLine[0], " ")
		output := utils.SpliceBy(splitLine[1], " ")

		lines = append(lines, line{
			signal: *(*[10]string)(signal),
			output: *(*[4]string)(output),
		})
	}
	return lines
}

func part1(lines []line) int {
	var sum int
	for _, line := range lines {
		for _, s := range line.output {
			switch len(s) {
			case 2:
				//1
				sum++
			case 3:
				//7
				sum++
			case 4:
				//4
				sum++
			case 7:
				//8
				sum++
			}
		}
	}
	return sum
}

func part2(lines []line) int {
	correctMap := workOutCorrectValues()

	var sum int
	for _, l := range lines {
		value, err := strconv.Atoi(workOutOutput(l, correctMap))
		if err != nil {
			panic("Not a number")
		}
		sum += value
	}
	return sum
}

func workOutCorrectValues() map[int]string {
	//numbers 0 1 2 3 4 5 6 7 8 9
	signal := [10]string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}

	countOfEachLetter := make(map[string]int)
	for _, s := range signal {
		for _, letter := range s {
			countOfEachLetter[string(letter)]++
		}
	}

	var sumOfLettersMakingUpEachSignal [10]int
	for i, s := range signal {
		for _, letter := range s {
			sumOfLettersMakingUpEachSignal[i] += countOfEachLetter[string(letter)]
		}
	}

	mapOfSumsToCorrectLetter := make(map[int]string)
	for i := 0; i <= 9; i++ {
		mapOfSumsToCorrectLetter[sumOfLettersMakingUpEachSignal[i]] = strconv.Itoa(i)
	}

	return mapOfSumsToCorrectLetter
}

func workOutOutput(line line, mapOfSumsToCorrectLetter map[int]string) string {

	countOfEachLetterInSignal := make(map[string]int)
	for _, s := range line.signal {
		for _, letter := range s {
			countOfEachLetterInSignal[string(letter)]++
		}
	}

	var actualOutput string
	for _, o := range line.output {
		var sumOfLettersInEachOutput int
		for _, letter := range o {
			sumOfLettersInEachOutput += countOfEachLetterInSignal[string(letter)]
		}
		actualOutput += mapOfSumsToCorrectLetter[sumOfLettersInEachOutput]
	}

	return actualOutput
}
