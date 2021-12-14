package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"math"
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

	inputs, rules := readLines(values)

	fmt.Printf("Part1- %v\n", part1(10, inputs, rules))
	fmt.Printf("Part2- %v\n", part1(40, inputs, rules))
}

func readLines(lines []string) (string, map[string]string) {
	input := lines[0]

	rules := make(map[string]string)

	for _, line := range lines[1:] {
		split := strings.Split(line, " -> ")
		rules[split[0]] = split[1]
	}
	return input, rules
}

func part1(numberOfLoops int, input string, rules map[string]string) int {
	counts := make(map[string]int)

	inputMap := make(map[string]int)
	for i, c := range input {
		if i < len(input)-1 {
			inputPair := string(input[i]) + string(input[i+1])
			if _, exists := inputMap[inputPair]; exists {
				inputMap[inputPair]++
			} else {
				inputMap[inputPair] = 1
			}
		}

		char := string(c)
		if _, exists := counts[char]; exists {
			counts[char]++
		} else {
			counts[char] = 1
		}
	}

	for i := 0; i < numberOfLoops; i++ {
		newInputMap, newCount := processRules(rules, counts, inputMap)
		inputMap = newInputMap
		counts = newCount
	}

	min := math.MaxInt
	max := 0
	for _, count := range counts {
		if min > count {
			min = count
		}
		if max < count {
			max = count
		}
	}
	return max - min
}

func processRules(rules map[string]string, counts map[string]int, inputMap map[string]int) (map[string]int, map[string]int) {
	newMap := make(map[string]int)

	newCounts := make(map[string]int)
	for key, value := range counts {
		newCounts[key] = value
	}

	for input, currentCount := range inputMap {
		toLetter := rules[input]

		newInputValue1 := input[:1] + toLetter
		newInputValue2 := toLetter + input[1:]
		if _, exists := newMap[newInputValue1]; exists {
			newMap[newInputValue1] += currentCount
		} else {
			newMap[newInputValue1] = currentCount
		}
		if _, exists := newMap[newInputValue2]; exists {
			newMap[newInputValue2] += currentCount
		} else {
			newMap[newInputValue2] = currentCount
		}

		//update the counts
		if _, exists := newCounts[toLetter]; exists {
			newCounts[toLetter] += currentCount
		} else {
			newCounts[toLetter] = currentCount
		}
	}

	return newMap, newCounts
}
