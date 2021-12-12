package main

import (
	"fmt"
	"github.com/jasonquinn/aoc2021/utils"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	file := string(input)
	values := utils.SpliceByLine(file)

	fmt.Printf("Part1- %v\n", part1(values))
	fmt.Printf("Part2- %v\n", part2(values))
}

func part1(lines []string) int {
	_, values := findPossiblePaths(createConnections(lines), "", 0)
	return len(values)
}

func part2(lines []string) int {
	connections := createConnections(lines)

	foundPaths := make(map[string]bool)
	for key := range connections.Value {
		if utils.IsLower(key) && key != "start" && key != "end" {
			_, values := findPossiblePaths(connections, key, 1)
			for _, value := range values {
				foundPaths[value] = true
			}
		}
	}
	return len(foundPaths)
}

func findPossiblePaths(connections utils.MultiMap, allowedDuplicate string, allowedDuplicateCount int) (bool, []string) {
	visited := make(map[string]bool)
	visited["start"] = true
	return hasPossiblePaths(connections, "start", "", visited, allowedDuplicate, allowedDuplicateCount)
}

func hasPossiblePaths(connections utils.MultiMap, startPosition string, currentPath string, visitedBefore map[string]bool, allowedDuplicate string, allowedDuplicateCount int) (bool, []string) {
	var possiblePaths []string
	if startPosition == "end" {
		return true, append(possiblePaths, currentPath+","+startPosition)
	}
	connectedPaths := connections.GetValues(startPosition)
	if connectedPaths == nil {
		return false, nil
	}
	anyPossible := false

	for connectedPath := range connectedPaths {
		//check if got before
		if utils.IsUpper(connectedPath) || visitedBefore[connectedPath] == false || (connectedPath == allowedDuplicate && allowedDuplicateCount >= 0) {
			newVisitedBefore := copyMap(visitedBefore)
			newVisitedBefore[connectedPath] = true

			var hasPossible bool
			var currentPossiblePaths []string

			if allowedDuplicate == connectedPath {
				hasPossible, currentPossiblePaths = hasPossiblePaths(connections, connectedPath, currentPath+","+startPosition, newVisitedBefore, allowedDuplicate, allowedDuplicateCount-1)
			} else {
				hasPossible, currentPossiblePaths = hasPossiblePaths(connections, connectedPath, currentPath+","+startPosition, newVisitedBefore, allowedDuplicate, allowedDuplicateCount)
			}

			if hasPossible {
				for _, path := range currentPossiblePaths {
					possiblePaths = append(possiblePaths, path)
				}
				anyPossible = true
			}
		}
	}
	if anyPossible {
		return true, possiblePaths
	}
	return false, nil
}

func createConnections(connections []string) utils.MultiMap {
	connectionMap := utils.NewMultiMap()

	for _, connection := range connections {
		splitConnectiokn := utils.SpliceBy(connection, "-")

		connectionMap.AddValue(splitConnectiokn[0], splitConnectiokn[1])
		connectionMap.AddValue(splitConnectiokn[1], splitConnectiokn[0])
	}
	return connectionMap
}

func copyMap(originalMap map[string]bool) map[string]bool {
	// Create the target map
	targetMap := make(map[string]bool)

	// Copy from the original map to the target map
	for key, value := range originalMap {
		targetMap[key] = value
	}
	return targetMap
}
