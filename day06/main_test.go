package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLanternFish(t *testing.T) {
	assert.Equal(t, 2, multipleLanternFish([]int{3}, 5))
}

func TestMultipleLanternFish(t *testing.T) {
	assert.Equal(t, 26, multipleLanternFish([]int{3, 4, 3, 1, 2}, 18))
	assert.Equal(t, 5934, multipleLanternFish([]int{3, 4, 3, 1, 2}, 80))
}
