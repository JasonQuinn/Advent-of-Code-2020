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

	{
		forward, depth, err := CalculatePosition(values)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("Part1: %v\n", forward*depth)
	}
	{
		forward, depth, err := CalculatePositionWithAim(values)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("Part2: %v\n", forward*depth)
	}
}

func CalculatePosition(values []string) (int, int, error) {
	forward := 0
	depth := 0
	for _, value := range values {
		splitString := strings.Split(value, " ")
		direction := splitString[0]
		distance, err := strconv.Atoi(splitString[1])
		if err != nil {
			return 0, 0, err
		}

		switch direction {
		case "forward":
			forward += distance
		case "down":
			depth += distance
		case "up":
			depth -= distance

		}
	}
	return forward, depth, nil
}

func CalculatePositionWithAim(values []string) (int, int, error) {
	aim := 0
	forward := 0
	depth := 0
	for _, value := range values {
		splitString := strings.Split(value, " ")
		direction := splitString[0]
		distance, err := strconv.Atoi(splitString[1])
		if err != nil {
			return 0, 0, err
		}

		switch direction {
		case "forward":
			forward += distance
			depth += distance * aim
		case "down":
			aim += distance
		case "up":
			aim -= distance

		}
	}
	return forward, depth, nil
}
