package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values, err := utils.SpliceByToInt(file, ",")
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Part1- %v\n", part1(values))
	fmt.Printf("Part2- %v\n", part2(values))
}

func part1(crabs []int) int {
	min, max := minMaxInArray(crabs)

	var minFuel int
	for i := min; i < max; i++ {

		var requiredFuel int
		for _, crab := range crabs {
			if i < crab {
				requiredFuel += crab - i
			} else {
				requiredFuel += i - crab
			}
		}
		if minFuel == 0 || minFuel > requiredFuel {
			minFuel = requiredFuel
		}
	}
	return minFuel
}

func part2(crabs []int) int {
	min, max := minMaxInArray(crabs)

	var minFuel int
	for i := min; i < max; i++ {

		var requiredFuel int
		for _, crab := range crabs {
			difference := 0
			if i < crab {
				difference += crab - i
			} else {
				difference += i - crab
			}
			requiredFuel += (difference * (difference + 1)) / 2

		}
		if minFuel == 0 || minFuel > requiredFuel {
			minFuel = requiredFuel
		}
	}
	return minFuel
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func minMaxInArray(array []int) (int, int) {
	min, max := 0, 0
	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
