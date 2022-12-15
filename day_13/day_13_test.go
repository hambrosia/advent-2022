package main

import (
	"testing"

	. "github.com/hambrosia/advent-2022/helpers"
)

func TestD13(t *testing.T) {
	p1SmallInput := GetInput("small_input.txt")
	gotP1Small := SumPacketsInOrder(p1SmallInput, true)
	wantP1Small := 13
	AssertEquals(t, gotP1Small, wantP1Small)

}
