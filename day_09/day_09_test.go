package main

import (
	"testing"
)

func TestRope(t *testing.T) {

	// P1 Small
	smallMotions := GetData("small_input.txt")
	p1SmallInputRope := MakeRope(2)
	p1SmallInputRope.Move(smallMotions)
	gotSmallHistory := len(p1SmallInputRope.history)
	wantSmallHistory := 13
	if gotSmallHistory != wantSmallHistory {
		t.Errorf("small rope, small input got %v, want %v", gotSmallHistory, wantSmallHistory)
	}

	// P1 Large
	largeMotions := GetData("large_input.txt")
	p1LargeInputRope := MakeRope(2)
	p1LargeInputRope.Move(largeMotions)
	gotLargeHistory := len(p1LargeInputRope.history)
	wantLargeHistory := 6503
	if gotLargeHistory != wantLargeHistory {
		t.Errorf("small rope, big input got %v, want %v", gotLargeHistory, wantLargeHistory)
	}

	// P2 Small
	smallMotionsP2 := GetData("small_input_p2.txt")
	p2SmallInputRope := MakeRope(10)
	p2SmallInputRope.Move(smallMotionsP2)
	p2GotSmallHistory := len(p2SmallInputRope.history)
	p2WantSmallHistory := 36
	if p2GotSmallHistory != p2WantSmallHistory {
		t.Errorf("big rope, small input got %v, want %v", p2GotSmallHistory, p2WantSmallHistory)
	}

	// P2 Large
	bigMotionsP2 := GetData("large_input.txt")
	p2LargeInputRope := MakeRope(10)
	p2LargeInputRope.Move(bigMotionsP2)
	p2GotBigHistory := len(p2LargeInputRope.history)
	p2WantBigHistory := 2724
	if p2GotBigHistory != p2WantBigHistory {
		t.Errorf("big rope, big input got %v, want %v", p2GotBigHistory, p2WantBigHistory)
	}

}
