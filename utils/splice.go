package utils

import (
	"strconv"
	"strings"
)

func SpliceByLineToInts(input string) ([]int, error) {
	var ints []int
	stringValues, err := SpliceByLine(input)

	if err != nil {
		return nil, err
	}

	for _, line := range stringValues {
		lineValue, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ints = append(ints, lineValue)
	}
	return ints, nil
}

func SpliceByLine(input string) ([]string, error) {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n"), nil
}
