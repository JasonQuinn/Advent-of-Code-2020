package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
)

func main() {
	input, err := os.ReadFile("C:\\Users\\Jason\\go\\src\\Advent-of-Code-2021\\01\\input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values, err := utils.InputToIntSplice(file)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", CalculateNumOfIncreases(values))
	fmt.Printf("Part 2: %v\n", CalculateNumOfSlidingIncrease(values))
}

func CalculateNumOfIncreases(values []int) int {
	increments := 0
	var oldValue int
	for i, value := range values {
		if i != 0 {
			if oldValue < value {
				increments++
			}
		}
		oldValue = value
	}
	return increments
}

func CalculateNumOfSlidingIncrease(values []int) int {
	increments := 0
	var previousValue int = -1
	for i := 0; i < len(values)-2; i++ {
		newValue := values[i] + values[i+1] + values[i+2]
		if previousValue != -1 && previousValue < newValue {
			increments++
		}
		previousValue = newValue
	}
	return increments
}
