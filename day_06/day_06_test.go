package day06

import (
	"fmt"
	"testing"
)

func TestDay06(t *testing.T) {

	testInput, _ := GetTestData("large_input.txt")
	fmt.Println(testInput)

	want := 1896
	got := GetStartOfFirstPacket(testInput, 3)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	wantMessage := 3452
	gotMessage := GetStartOfFirstPacket(testInput, 13)
	if gotMessage != wantMessage {
		t.Errorf("got %v, want %v", gotMessage, wantMessage)
	}

}
