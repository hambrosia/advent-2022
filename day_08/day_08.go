package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetTestData(filename string) (testData [][]int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read test input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	testData = make([][]int, 0)

	for scanner.Scan() {
		treeLine := scanner.Text()
		treeFields := strings.Split(treeLine, "")
		treeInts := make([]int, 0)
		for _, t := range treeFields {
			tInt, _ := strconv.Atoi(t)
			treeInts = append(treeInts, tInt)
		}
		testData = append(testData, treeInts)
	}

	return testData, err
}

func IsVisible(index int, val int, view []int, name string) bool {
	if index == 0 || index == len(view)-1 {
		return true
	}

	leftViz, rightViz := true, true
	for _, compareVal := range view[:index] {
		if compareVal >= val {
			leftViz = false
		}
	}

	for _, compareVal := range view[index+1:] {
		if compareVal >= val {
			rightViz = false
		}
	}

	ret := (leftViz || rightViz)
	return ret
}

func CalculateScenicScore(tree int, view []int, name string) int {
	var i int
	treeSeen := false
	for i = 0; i < len(view); i++ {
		if view[i] > 0 {
			treeSeen = true
		}
		if view[i] >= tree {
			i = i + 1
			break
		}
	}
	if !treeSeen {
		return 0
	}
	ret := i
	return ret
}

func FindVisibleTrees(trees [][]int) (visibleCount int, visibleMap map[string]bool) {
	visibleMap = map[string]bool{}
	for i := 0; i < len(trees); i++ {
		row := trees[i]
		for j := 0; j < len(row); j++ {
			// check if visible

			val := trees[i][j]
			xView := trees[i]
			yView := make([]int, 0)

			for k := 0; k < len(trees); k++ {
				yView = append(yView, trees[k][j])
			}

			yViz := IsVisible(i, val, yView, "yView")
			xViz := IsVisible(j, val, xView, "xView")
			isViz := yViz || xViz

			if isViz {
				name := strconv.Itoa(i) + strconv.Itoa(j)
				visibleMap[name] = true
				visibleCount++
			}

		}
	}
	return visibleCount, visibleMap
}

func ReverseIntSlice(original []int) (reversed []int) {
	reversed = make([]int, 0)
	for i := len(original) - 1; i >= 0; i-- {
		reversed = append(reversed, original[i])
	}
	return reversed
}

func FindMostScenic(trees [][]int) (mostScenicValue int) {
	var mostScenicTree []int
	for i := 0; i < len(trees); i++ {
		row := trees[i]
		for j := 0; j < len(row); j++ {

			height := trees[i][j]
			if height < 1 {
				continue
			}

			xView := trees[i]
			yView := make([]int, 0)

			for k := 0; k < len(trees); k++ {
				yView = append(yView, trees[k][j])
			}

			right := xView[j+1:]
			down := yView[i+1:]

			var leftReverse = ReverseIntSlice(xView[:j])
			var upReverse = ReverseIntSlice(yView[:i])

			score := CalculateScenicScore(height, right, "right") * CalculateScenicScore(height, leftReverse, "left") * CalculateScenicScore(height, down, "down") * CalculateScenicScore(height, upReverse, "up")
			if score > mostScenicValue {
				mostScenicValue = score
				mostScenicTree = []int{i, j}
			}
		}
	}
	fmt.Println("most scenic tree", mostScenicTree)
	fmt.Println("most scenic score", mostScenicValue)
	return mostScenicValue
}
