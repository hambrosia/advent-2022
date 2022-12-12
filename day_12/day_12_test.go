package main

import (
	"testing"

	. "github.com/hambrosia/advent-2022/helpers"
)

func TestD12(t *testing.T) {
	heightmap, start, end := GetInput("small_input.txt")
	gotSteps, gotPath := FindShortestPath(heightmap, start, end)
	PrintPath(heightmap, gotPath)
	AssertEquals(t, gotSteps, 31)

	// heightmap, start, end = GetInput("large_input.txt")
	// got, _ := FindShortestPath(heightmap, start, end)
	// AssertEquals(t, got, 31)

}
