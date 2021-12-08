package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	lineSegments := readLines([]string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	})

	assert.Equal(t, 5, part1(lineSegments))
}

func TestPart2(t *testing.T) {
	lineSegments := readLines([]string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	})

	assert.Equal(t, 12, part2(lineSegments))
}

func TestLineSegments(t *testing.T) {
	lineSegments := readLines([]string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	})

	assert.Equal(t, 10, len(lineSegments))

	assert.Equal(t, 0, lineSegments[0].start.x)
	assert.Equal(t, 9, lineSegments[0].start.y)
	assert.Equal(t, 5, lineSegments[0].end.x)
	assert.Equal(t, 9, lineSegments[0].end.y)

	assert.Equal(t, 8, lineSegments[1].start.x)
	assert.Equal(t, 0, lineSegments[1].start.y)
	assert.Equal(t, 0, lineSegments[1].end.x)
	assert.Equal(t, 8, lineSegments[1].end.y)
}

func TestNumbersBetween(t *testing.T) {
	{
		points := pointsBetween(lineSegment{start: point{x: 1, y: 1}, end: point{
			x: 1,
			y: 3,
		}}, true)
		assert.Equal(t, 3, len(points))
		assert.Equal(t, 1, points[1].x)
		assert.Equal(t, 2, points[1].y)
	}
	{
		points := pointsBetween(lineSegment{start: point{x: 9, y: 7}, end: point{
			x: 7,
			y: 7,
		}}, false)
		assert.Equal(t, 3, len(points))
		assert.Equal(t, 8, points[1].x)
		assert.Equal(t, 7, points[1].y)
	}
	{
		points := pointsBetween(lineSegment{start: point{x: 1, y: 1}, end: point{
			x: 3,
			y: 3,
		}}, false)
		assert.Equal(t, 3, len(points))
		assert.Equal(t, 2, points[1].x)
		assert.Equal(t, 2, points[1].y)
	}
	{
		points := pointsBetween(lineSegment{start: point{x: 9, y: 7}, end: point{
			x: 7,
			y: 9,
		}}, false)
		assert.Equal(t, 3, len(points))
		assert.Equal(t, 8, points[1].x)
		assert.Equal(t, 8, points[1].y)
	}
	{
		points := pointsBetween(lineSegment{start: point{x: 9, y: 7}, end: point{
			x: 7,
			y: 9,
		}}, true)
		assert.Equal(t, 0, len(points))
	}
}
