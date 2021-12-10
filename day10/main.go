package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
	"sort"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values := utils.SpliceByLine(file)

	fmt.Printf("Part1- %v\n", part1(values))
	fmt.Printf("Part2- %v\n", part2(values))
}

func part2(input []string) int {
	var scores []int

	for _, line := range input {
		var validSoFar []string
		lineIsValid := true
		for _, c := range line {
			char := string(c)

			if char == "(" || char == "[" || char == "{" || char == "<" {
				validSoFar = append(validSoFar, char)
			} else {
				if char == ")" || char == "]" || char == "}" || char == ">" {
					wasCorrect, newValid := ifExpectedRemove(validSoFar, getOpposite(char))
					if wasCorrect {
						validSoFar = newValid
					} else {
						lineIsValid = false
						break
					}
				}
			}
		}
		if lineIsValid {
			var sum int
			for i := len(validSoFar) - 1; i >= 0; i-- {
				sum *= 5
				sum += getValuePart2(getOpposite(validSoFar[i]))
			}
			scores = append(scores, sum)
			fmt.Printf("Line - %v score- %v\n", validSoFar, sum)
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})
	return scores[len(scores)/2]
}

func part1(input []string) int {
	var sum int
	for _, line := range input {
		var validSoFar []string
		for _, c := range line {
			char := string(c)

			if char == "(" || char == "[" || char == "{" || char == "<" {
				validSoFar = append(validSoFar, char)
			} else {
				if char == ")" || char == "]" || char == "}" || char == ">" {
					wasCorrect, newValid := ifExpectedRemove(validSoFar, getOpposite(char))
					if wasCorrect {
						validSoFar = newValid
					} else {
						sum += getValuePart1(char)
						break
					}
				}

			}
		}
	}
	return sum
}

func ifExpectedRemove(array []string, expected string) (bool, []string) {
	if array[len(array)-1] == expected {
		return true, array[:len(array)-1]
	} else {
		return false, array
	}
}

func getOpposite(char string) string {
	charMap := make(map[string]string)
	charMap[")"] = "("
	charMap["]"] = "["
	charMap["}"] = "{"
	charMap[">"] = "<"
	charMap["("] = ")"
	charMap["["] = "]"
	charMap["{"] = "}"
	charMap["<"] = ">"

	return charMap[char]
}

func getValuePart1(char string) int {
	charMap := make(map[string]int)
	charMap[")"] = 3
	charMap["]"] = 57
	charMap["}"] = 1197
	charMap[">"] = 25137

	return charMap[char]
}

func getValuePart2(char string) int {
	charMap := make(map[string]int)
	charMap[")"] = 1
	charMap["]"] = 2
	charMap["}"] = 3
	charMap[">"] = 4

	return charMap[char]
}
