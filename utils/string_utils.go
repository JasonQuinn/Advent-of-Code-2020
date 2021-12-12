package utils

import (
	"strconv"
	"unicode"
)

func BinaryStringToInt(binary string) int {
	value, err := strconv.ParseInt(binary, 2, 32)
	if err != nil {
		panic("Not an binary")
	}

	return int(value)
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
