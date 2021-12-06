package main

import "testing"

func TestPart1(t *testing.T) {
	values := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	got := CalculateNumOfIncreases(values)
	if got != 7 {
		t.Errorf("Error")
	}
}

func TestPart2(t *testing.T) {
	values := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	got := CalculateNumOfSlidingIncrease(values)
	if got != 5 {
		t.Errorf("Error")
	}
}
