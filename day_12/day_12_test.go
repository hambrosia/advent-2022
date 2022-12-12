package main

import (
	"fmt"
	"testing"

	. "github.com/hambrosia/advent-2022/helpers"
)

func TestD12(t *testing.T) {
	heightmap, start, end := GetInput("small_input.txt")
	gotSteps, gotPath := FindShortestPath(heightmap, start, end)
	PrintPath(heightmap, gotPath)
	AssertEquals(t, gotSteps, 31)
	for i, p := range gotPath {
		fmt.Println("step", i)
		fmt.Println("p", p)
		fmt.Println(heightmap[p.y][p.x])
		fmt.Println()
	}

	heightmap, start, end = GetInput("large_input.txt")
	gotLarge, _ := FindShortestPath(heightmap, start, end)
	AssertEquals(t, gotLarge, 447)

	heightmap, _, end := GetInput("small_input.txt")
	elevation := 0
	gotSteps, gotPath := FindShortestPathFromElevation(heightmap, elevation, end)
	PrintPath(heightmap, gotPath)
	AssertEquals(t, gotSteps, 31)

}
