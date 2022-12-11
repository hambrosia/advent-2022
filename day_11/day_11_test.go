package main

import (
	"math/big"
	"testing"
)

func TestMonkeyBusiness(t *testing.T) {

	// p1 small
	monkeysP1Small := makeMonkeys("small_input.txt")
	lombP1Small := DoRounds(monkeysP1Small, 20, 3, false)
	wantP1Small := big.NewInt(10605)
	if lombP1Small.Cmp(wantP1Small) != 0 {
		t.Errorf("p1 small input, got %v, want %v", lombP1Small, wantP1Small)
	}

	// p1 large
	monkeyP1Large := makeMonkeys("large_input.txt")
	lombP1Large := DoRounds(monkeyP1Large, 20, 3, false)
	wantP1Large := big.NewInt(119715)
	if lombP1Large.Cmp(wantP1Large) != 0 {
		t.Errorf("p1 large input, got %v, want %v", lombP1Large, wantP1Large)
	}

	// p2 small
	monkeyP2Small := makeMonkeys("small_input.txt")
	lombP2Small := DoRounds(monkeyP2Small, 10000, 1, true)
	wantP2Small := big.NewInt(2713310158)
	if lombP2Small.Cmp(wantP2Small) != 0 {
		t.Errorf("p2 small input got %v, want %v", lombP2Small, wantP2Small)
	}

	// 2637590098 too low for large input

}
