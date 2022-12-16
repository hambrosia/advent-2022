package main

import (
	"testing"

	. "github.com/hambrosia/advent-2022/helpers"
)

func TestD13(t *testing.T) {
	p1SmallInput := GetInput("small_input.txt")
	gotP1Small := SumPacketsInOrder(p1SmallInput)
	wantP1Small := 13
	AssertEquals(t, gotP1Small, wantP1Small)

	p1LargeInput := GetInput("large_input.txt")
	gotP1Large := SumPacketsInOrder(p1LargeInput)
	wantP1Large := -1
	AssertEquals(t, gotP1Large, wantP1Large)

}
