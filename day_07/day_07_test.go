package main

import (
	"testing"
)

func TestDay07(t *testing.T) {

	testSmallInput, err := GetTestData("small_input.txt")
	if err != nil {
		t.Error("error reading file", err)
	}
	tree, sizes := ParseTestData(testSmallInput)
	p1Got, p2Got := GetSizeOfFolders(tree, sizes)

	p1Want := 95437
	if p1Got != p1Want {
		t.Errorf("got %v, want %v", p1Got, p1Want)
	}

	p2Want := 24933642

	if p2Got != p2Want {
		t.Errorf("got %v, want %v", p2Got, p2Want)
	}

	testBigInput, errBig := GetTestData("large_input.txt")
	if errBig != nil {
		t.Error("error reading file", err)
	}
	treeBig, sizesBig := ParseTestData(testBigInput)
	p1GotBig, p2GotBig := GetSizeOfFolders(treeBig, sizesBig)

	p1WantBig := 1611443
	if p1GotBig != p1WantBig {
		t.Errorf("got %v, want %v", p1GotBig, p1WantBig)
	}

	p2WantBig := -1
	maxFileSystem := 70000000
	if p2GotBig > maxFileSystem {
		t.Errorf("got %v, want < %v", p2GotBig, maxFileSystem)
	}
	if p2GotBig != p1WantBig {
		t.Errorf("got %v, want %v", p2GotBig, p2WantBig)
	}

}
