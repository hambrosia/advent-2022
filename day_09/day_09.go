package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetData(filename string) (motions []Motion) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("unable to read test data")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	motions = make([]Motion, 0)

	directionKey := map[string]Pair{
		"R": {x: 1, y: 0},
		"L": {x: -1, y: 0},
		"U": {x: 0, y: 1},
		"D": {x: 0, y: -1},
	}

	for scanner.Scan() {
		rawText := scanner.Text()
		textFields := strings.Fields(rawText)
		val, _ := strconv.Atoi(textFields[1])
		newMotion := Motion{textFields[0], directionKey[textFields[0]], val}
		motions = append(motions, newMotion)
	}
	return motions
}

type Pair struct {
	x, y int
}

type Motion struct {
	directionName string
	direction     Pair
	steps         int
}

func MoveH(coordinates *Pair, motion Motion) {
	coordinates.x += motion.direction.x
	coordinates.y += motion.direction.y
}

func MoveT(h Pair, t Pair, tHistory *map[Pair]struct{}) {
	// move t based on relation to h
	// if t moved, upsert t location into history
}

func main() {
	// rope behavior
	// 1 the head (H) and tail (T) must always be touching (diagonally adjacent and even overlapping both count as touching
	// If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move one step in that direction so it remains close enough
	// Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves one step diagonally to keep up
	// output is number of positions (T) occupied at least once, e.g. set of T coordinates

	motions := GetData("small_input.txt")
	head := Pair{0, 0}
	tail := Pair{0, 0}
	tHistory := make(map[Pair]struct{})
	fmt.Println("head", head)
	for _, motion := range motions {
		for range make([]int, motion.steps) {
			MoveH(&head, motion)
			MoveT(head, tail, &tHistory)
			fmt.Println("head", head)
		}
	}

}
