package utils

import (
	"strconv"
	"strings"
)

func SpliceByLineToInts(input string) ([]int, error) {
	return SpliceByToInt(input, "\n")
}

func SpliceByLine(input string) []string {
	return SpliceBy(input, "\n")
}

func SpliceByToInt(input string, sep string) ([]int, error) {
	var ints []int
	stringValues := SpliceBy(input, sep)

	for _, line := range stringValues {
		lineValue, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ints = append(ints, lineValue)
	}
	return ints, nil
}

func SpliceBy(input string, sep string) []string {
	spliced := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), sep)

	var filtered []string
	for i := range spliced {
		if spliced[i] != " " && spliced[i] != "" {
			filtered = append(filtered, spliced[i])
		}
	}

	return filtered
}
