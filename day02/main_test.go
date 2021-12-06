package main

import "testing"

func TestPart1(t *testing.T) {
	values := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	forward, depth, err := CalculatePosition(values)
	if err != nil {
		t.Errorf("Error")
	}
	if forward != 15 {
		t.Errorf("Error")
	}
	if depth != 10 {
		t.Errorf("Error")
	}
}

func TestPart2(t *testing.T) {
	values := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	forward, depth, err := CalculatePositionWithAim(values)
	if err != nil {
		t.Errorf("Error")
	}
	if forward != 15 {
		t.Errorf("Error")
	}
	if depth != 60 {
		t.Errorf("Error")
	}
}
