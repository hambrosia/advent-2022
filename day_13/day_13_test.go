package main

import (
	"testing"
)

func TestD13(t *testing.T) {
	p1SmallInput := GetInput("large_input.txt")
	ComparePackets(p1SmallInput, true)
	// wantP1Small := 13
}
