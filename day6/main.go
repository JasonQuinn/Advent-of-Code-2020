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

	fmt.Printf("Part1- %v", multipleLanternFish(values, 80))
	fmt.Printf("Part2- %v", multipleLanternFish(values, 256))
}

func multipleLanternFish(startPeriods []int, daysToRun int) int {

	var days [9]int
	for _, period := range startPeriods {
		days[period]++
	}

	for day := 0; day < daysToRun; day++ {
		//var newDays [8]int
		givingBirth := days[0]
		for i := 0; i < 8; i++ {
			days[i] = days[i+1]
		}
		days[6] += givingBirth
		days[8] = givingBirth
		fmt.Printf("Day %v, Values- %v\n", day, days)
	}

	var sum int
	for _, day := range days {
		sum += day
	}
	return sum
}
