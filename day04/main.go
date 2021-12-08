package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
)

type Bingo struct {
	Numbers []int
	Cards   [][][]int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values := utils.SpliceByLine(file)

	bingo, _ := getNumbersAndCards(values)

	fmt.Printf("Part1: %v\n", part1(bingo))
	fmt.Printf("Part2: %v\n", part2(bingo))
}

func part1(bingo Bingo) int {
	for _, number := range bingo.Numbers {
		for _, card := range bingo.Cards {
			for _, line := range card {
				for i, value := range line {
					//filter out matching number
					if value == number {
						line[i] = -1
					}
				}
			}
			
			matched, sum := checkMatch(card)
			if matched {
				return sum * number
			}
		}
	}
	return 0
}

func part2(bingo Bingo) int {
	foundCardsCount := 0
	for _, number := range bingo.Numbers {
		for j, card := range bingo.Cards {
			if bingo.Cards[j] != nil {

				for _, line := range card {
					for i, value := range line {
						//filter out matching number
						if value == number {
							line[i] = -1
						}
					}
				}

				matched, sum := checkMatch(card)
				if matched {
					bingo.Cards[j] = nil
					foundCardsCount++

					if foundCardsCount == len(bingo.Cards) {
						return sum * number
					}
				}
			}
		}
	}
	return 0
}

func checkMatch(card [][]int) (bool, int) {
	//check if matching rows
	for _, line := range card {
		allEmpty := true
		for _, value := range line {
			if value != -1 {
				allEmpty = false
			}
		}
		if allEmpty {
			sum := 0
			for _, matchingLine := range card {
				for _, matchingValue := range matchingLine {
					if matchingValue != -1 {
						sum += matchingValue
					}
				}
			}
			return true, sum
		}
	}
	//check column
	for i := 0; i < 5; i++ {
		allEmpty := true
		for _, line := range card {
			if line[i] != -1 {
				allEmpty = false
			}
		}
		if allEmpty {
			sum := 0
			for _, matchingLine := range card {
				for _, matchingValue := range matchingLine {
					if matchingValue != -1 {
						sum += matchingValue
					}
				}
			}

			return true, sum
		}
	}
	return false, -1
}

func getNumbersAndCards(values []string) (Bingo, error) {
	var bingo Bingo

	var currentCard [][]int
	for i, value := range values {
		if i == 0 {
			numbers, err := utils.SpliceByToInt(value, ",")
			if err != nil {
				return bingo, err
			}
			bingo.Numbers = numbers
		} else if value != "" {
			line, err := utils.SpliceByToInt(value, " ")
			if err != nil {
				return bingo, err
			}
			currentCard = append(currentCard, line)
		}
		if len(currentCard) == 5 {
			bingo.Cards = append(bingo.Cards, currentCard)
			currentCard = nil
		}
	}
	return bingo, nil
}
