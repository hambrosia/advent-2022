package main

import (
	"testing"

	. "github.com/hambrosia/advent-2022/helpers"
)

func TestD12(t *testing.T) {

	// p1 small
	heightmap, start, end := GetInput("small_input.txt")
	gotSteps, _ := FindShortestPath(heightmap, start, end, false)
	AssertEquals(t, gotSteps, 31)

	// p1 big
	heightmap, start, end = GetInput("large_input.txt")
	gotLarge, _ := FindShortestPath(heightmap, start, end, false)
	AssertEquals(t, gotLarge, 447)

	// p2 small
	heightmap2, _, end2 := GetInput("small_input.txt")
	elevation := 0
	gotSteps2 := FindShortestPathFromElevation(heightmap2, elevation, end2, false)
	AssertEquals(t, gotSteps2, 29)

	//p2 big
	heightmap3, _, end3 := GetInput("large_input.txt")
	gotSteps3 := FindShortestPathFromElevation(heightmap3, elevation, end3, false)
	AssertEquals(t, gotSteps3, 446)

}
