package main

import (
	"fmt"
	"testing"
)

func TestDay07(t *testing.T) {

	testInput, _ := GetTestData("large_input.txt")
	// fmt.Println(testInput)
	tree, sizes := parseTestData(testInput)
	res := getSizeOfFolders(tree, sizes)

	fmt.Println("Answer")
	fmt.Println(res)
}

