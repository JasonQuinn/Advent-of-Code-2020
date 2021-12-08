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
	values := utils.SpliceByLine(file)

	fmt.Printf("Part1: %v\n", CalculatePowerConsumption(values))
	fmt.Printf("Part2: %v\n", CalculateScrubber(values))
}

func CalculatePowerConsumption(values []string) int {
	lengthOfNumber := len(values[0])

	newNumber := make([]int, lengthOfNumber)
	for _, value := range values {
		for i := 0; i < lengthOfNumber; i++ {
			if string(value[i]) == "1" {
				newNumber[i]++
			}
		}
	}

	gamma := ""
	epsilon := ""
	for _, number := range newNumber {
		if number > len(values)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	return utils.BinaryStringToInt(gamma) * utils.BinaryStringToInt(epsilon)
}

func CalculateScrubber(values []string) int {
	return calculateO2(values) * calculateCo2(values)
}

func calculateO2(values []string) int {
	return calculateScrubber(values, 0, true)
}

func calculateCo2(values []string) int {
	return calculateScrubber(values, 0, false)
}

func calculateScrubber(values []string, currentPosition int, getO2 bool) int {
	if len(values) == 1 {
		return utils.BinaryStringToInt(values[0])
	}

	lengthOfNumber := len(values[0])
	if currentPosition >= lengthOfNumber {
		panic("Shouldnt happen")
	}

	newNumber := 0
	for _, value := range values {
		if string(value[currentPosition]) == "1" {
			newNumber++
		}
	}

	var prefix string
	if getO2 {
		if float32(newNumber) >= float32(len(values))/2 {
			prefix = "1"
		} else {
			prefix = "0"
		}
	} else {
		if float32(newNumber) >= float32(len(values))/2 {
			prefix = "0"
		} else {
			prefix = "1"
		}
	}

	var filtered []string
	for i := range values {
		if string(values[i][currentPosition]) == prefix {
			filtered = append(filtered, values[i])
		}
	}

	return calculateScrubber(filtered, currentPosition+1, getO2)
}
