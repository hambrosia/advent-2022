package main

import (
	"testing"
)

func TestD13(t *testing.T) {
	p1SmallInput := GetInput("small_input.txt")
	SumPacketsInOrder(p1SmallInput, true)
	// wantP1Small := 13
}
