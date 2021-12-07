package utils

import "strconv"

func BinaryStringToInt(binary string) int {
	value, err := strconv.ParseInt(binary, 2, 32)
	if err != nil {
		panic("Not an binary")
	}

	return int(value)
}
