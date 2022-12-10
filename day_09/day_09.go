package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Motion struct {
	directionName string
	direction     Pair
	steps         int
}

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

func (p *Pair) Push(motion Motion) {
	p.x += motion.direction.x
	p.y += motion.direction.y
}

func (p *Pair) Pull(h Pair, tHistory map[string]struct{}) {
	dist := p.Distance(h)
	if dist <= math.Sqrt(2) {
		// ensure position of tail in history and continue
		tHistory[p.ToString()] = struct{}{}
		return
	}
	aoi := p.AngleOfIncline(h)
	switch {
	case aoi == 0:
		p.x += 1
	case aoi == 90:
		p.y += 1
	case aoi == 180:
		p.x -= 1
	case aoi == 270:
		p.y -= 1
	case 0 < aoi && aoi < 90:
		p.y += 1
		p.x += 1
	case 90 < aoi && aoi < 180:
		p.y += 1
		p.x -= 1
	case 180 < aoi && aoi < 270:
		p.y -= 1
		p.x -= 1
	case 270 < aoi && aoi < 360:
		p.y -= 1
		p.x += 1
	}
	tHistory[p.ToString()] = struct{}{}
}

type Rope struct {
	length  int
	knots   []Pair
	history map[string]struct{}
}

func MakeRope(length int) Rope {
	return Rope{
		length,
		make([]Pair, length),
		map[string]struct{}{},
	}
}

func (rope *Rope) Move(motions []Motion) {
	for _, motion := range motions {
		for range make([]int, motion.steps) {
			for i := 0; i < rope.length; i++ {
				switch {
				case i == 0:
					rope.knots[i].Push(motion)
				case i == rope.length-1:
					rope.knots[i].Pull(rope.knots[i-1], rope.history)
				case 0 < i && i < rope.length-1:
					throwAwayHistory := make(map[string]struct{})
					rope.knots[i].Pull(rope.knots[i-1], throwAwayHistory)
				}
			}
		}
	}
}
