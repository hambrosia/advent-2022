package main

import (
	"bufio"
	"fmt"
	"math"
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

func (p1 Pair) AngleOfIncline(p2 Pair) (aoi float64) {
	aoi = math.Atan2(float64(p2.y)-float64(p1.y), float64(p2.x)-float64(p1.x)) * (180 / math.Pi)
	if aoi < 0 {
		aoi = 360 + aoi
	}
	return float64(aoi)
}

func (p1 Pair) Distance(p2 Pair) (distance float64) {
	xD := math.Pow(float64(p2.x-p1.x), 2)
	yD := math.Pow(float64(p2.y-p1.y), 2)
	distance = math.Sqrt(xD + yD)
	return distance
}

func (p Pair) ToString() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
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

func MoveT(h Pair, t *Pair, tHistory map[string]struct{}) {

	fmt.Println("START")
	fmt.Println("head", h)
	fmt.Println("tail", t)
	// move t based on relation to h
	dist := t.Distance(h)
	if dist <= math.Sqrt(2) {
		// ensure position of t in history and continue
		tHistory[t.ToString()] = struct{}{}
		return
	}
	fmt.Println("distance from t to h", dist)
	aoi := t.AngleOfIncline(h)
	fmt.Println("aoi from t to h", aoi)
	// move t one step in direction of h and update t history
	switch {
	case aoi == 0:
		t.x += 1
	case aoi == 90:
		t.y += 1
	case aoi == 180:
		t.x -= 1
	case aoi == 270:
		t.y -= 1
	case 0 < aoi && aoi < 90:
		t.y += 1
		t.x += 1
	case 90 < aoi && aoi < 180:
		t.y += 1
		t.x -= 1
	case 180 < aoi && aoi < 270:
		t.y -= 1
		t.x -= 1
	case 270 < aoi && aoi < 360:
		t.y -= 1
		t.x += 1
	}

	tHistory[t.ToString()] = struct{}{}
	fmt.Println("END")
	fmt.Println("head", h)
	fmt.Println("tail", t)

}

func main() {
	// rope behavior
	// the head (H) and tail (T) must always be touching (diagonally adjacent and even overlapping both count as touching
	// If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move one step in that direction so it remains close enough
	// Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves one step diagonally to keep up
	// output is number of positions (T) occupied at least once, e.g. set of T coordinates

	motions := GetData("large_input.txt")
	head := Pair{0, 0}
	tail := Pair{0, 0}
	tHistory := make(map[string]struct{})
	for _, motion := range motions {
		fmt.Println(motion.directionName, motion.steps)
		fmt.Println()
		for range make([]int, motion.steps) {
			MoveH(&head, motion)
			MoveT(head, &tail, tHistory)
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println("t history", tHistory)
	fmt.Println("history length", len(tHistory))

}
