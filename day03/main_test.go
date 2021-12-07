package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	values := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	assert.Equal(t, 198, CalculatePowerConsumption(values))
}

func TestPart2O2(t *testing.T) {
	values := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	assert.Equal(t, 23, calculateO2(values))
}

func TestPart2Co2(t *testing.T) {
	values := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	assert.Equal(t, 10, calculateCo2(values))
}

func TestPart2(t *testing.T) {
	values := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	assert.Equal(t, 230, CalculateScrubber(values))
}
