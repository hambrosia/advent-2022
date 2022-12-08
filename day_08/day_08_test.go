package main

import (
	"testing"
)

func TestTreeFinder(t *testing.T) {
	testDataSmall, errSmall := GetTestData("small_input.txt")
	if errSmall != nil {
		t.Error("Unable to read test data")
	}

	gotVisibleSmall, _ := FindVisibleTrees(testDataSmall)
	wantVisibleSmall := 21

	if gotVisibleSmall != wantVisibleSmall {
		t.Errorf("p1 small input got %v, want %v", gotVisibleSmall, wantVisibleSmall)
	}

	gotMostScenic := FindMostScenic(testDataSmall)
	wantMostScenic := 8
	if gotMostScenic != wantMostScenic {
		t.Errorf("p2 small most scenic got %v, want %v", gotMostScenic, wantMostScenic)
	}

	testDataBig, errBig := GetTestData("large_input.txt")
	if errBig != nil {
		t.Error("Unable to read test data")
	}

	gotVisibleBig, _ := FindVisibleTrees(testDataBig)
	wantVisibleBig := 1805

	if gotVisibleBig != wantVisibleBig {
		t.Errorf("p1 big input got %v, want %v", gotVisibleBig, wantVisibleBig)
	}

	gotMostScenicBig := FindMostScenic(testDataBig)
	wantMostScenicBig := 444528

	if gotMostScenicBig != wantMostScenicBig {
		t.Errorf("p2 big most scenic got %v, want %v", gotMostScenicBig, wantMostScenicBig)
	}

}
