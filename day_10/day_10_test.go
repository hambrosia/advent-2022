package main

import (
	"fmt"
	"testing"

	. "github.com/hambrosia/advent-2022/helpers"
)

func TestCRT(t *testing.T) {

	// p1 small
	instructions := GetInput("small_input.txt")
	sigStrength, _ := Process(instructions)
	AssertEquals(t, sigStrength, 13140)

	// p1 large
	instructions = GetInput("large_input.txt")
	sigStrength, _ = Process(instructions)
	AssertEquals(t, sigStrength, 17020)

	// p2 small
	instructions = GetInput("small_input.txt")
	_, displayOutput := Process(instructions)
	got := DisplayOutputToString(displayOutput)
	want := `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
	AssertEquals(t, got, want)

	// p2 large
	fmt.Println()
	instructions = GetInput("large_input.txt")
	_, displayOutput = Process(instructions)
	got = DisplayOutputToString(displayOutput)

	want = `###..#....####.####.####.#.....##..####.
#..#.#....#.......#.#....#....#..#.#....
#..#.#....###....#..###..#....#....###..
###..#....#.....#...#....#....#.##.#....
#.#..#....#....#....#....#....#..#.#....
#..#.####.####.####.#....####..###.####.`
	AssertEquals(t, got, want)

}
