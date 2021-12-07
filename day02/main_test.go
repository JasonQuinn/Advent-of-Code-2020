package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	values := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	forward, depth, err := CalculatePosition(values)
	assert.Nil(t, err)
	assert.Equal(t, 15, forward)
	assert.Equal(t, 10, depth)
}

func TestPart2(t *testing.T) {
	values := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	forward, depth, err := CalculatePositionWithAim(values)

	assert.Nil(t, err)
	assert.Equal(t, 15, forward)
	assert.Equal(t, 60, depth)
}
