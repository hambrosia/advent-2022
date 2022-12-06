package day04

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDay04(t *testing.T) {

	file, err := os.Open("input.txt")
	if err != nil {
		t.Fatal("Unable to read test input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([][]int, 0)
	for scanner.Scan() {
		shifts := strings.Split(scanner.Text(), ",")

		shift1 := strings.Split(shifts[0], "-")
		shift2 := strings.Split(shifts[1], "-")

		s1Start, _ := strconv.Atoi(shift1[0])
		s1End, _ := strconv.Atoi(shift1[1])
		s2Start, _ := strconv.Atoi(shift2[0])
		s2End, _ := strconv.Atoi(shift2[1])

		data = append(data, []int{s1Start, s1End, s2Start, s2End})
	}
	gotP1 := FindFullyOverlappingRegions(data)
	wantP1 := 431
	if gotP1 != wantP1 {
		t.Errorf("got %v, want %v", gotP1, wantP1)
	}
	gotP2 := FindPartiallyOverlappingRegions(data)
	wantP2 := 823
	if gotP2 != wantP2 {
		t.Errorf("got %v, want %v", gotP2, wantP2)
	}
}
