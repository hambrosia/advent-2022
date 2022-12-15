package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func GetInput(filename string) (ret [][]Point) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("bad rocks")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := []Point{}
		for _, pair := range strings.Split(line, "->") {
			values := strings.Split(strings.TrimSpace(pair), ",")
			x, _ := strconv.Atoi(values[0])
			y, _ := strconv.Atoi(values[1])
			coords = append(coords, Point{x, y})
		}
		ret = append(ret, coords)
	}
	return ret
}

func main() {
	input := GetInput("small_input.txt")
	for _, rock := range input {
		fmt.Println(rock)
	}
	maxX := 0
	maxY := 0
	minX := input[0][0].x
	minY := input[0][0].y
	for i := range input {
		for j := range input[i] {
			p := input[i][j]
			if p.x > maxX {
				maxX = p.x
			}
			if p.y > maxY {
				maxY = p.y
			}
			if p.x < minX {
				minX = p.x
			}
			if p.y < minY {
				minY = p.y
			}
		}
	}
	fmt.Println("max X", maxX, "maxY", maxY)
	fmt.Println("min X", minX, "minY", minY)
	cave := [200][600]string{}
	for _, line := range input {
		for i := 0; i < len(line)-1; i++ {
			p := line[i]
			n := line[i+1]
			// fmt.Println(p)
			// fmt.Println(n)
			// fmt.Println()
			if p.x == n.x {
				// draw vertical line
				start := p.y
				end := n.y
				direction := 0
				if start < end {
					direction = 1
				}
				if start > end {
					direction = -1
				}
				for ; start != end; start += direction {
					cave[start][p.x] = "#"
				}
			}
			if p.y == n.y {
				// draw horizontal line
				start := p.x
				end := n.x
				direction := 0
				if start < end {
					direction = 1
				}
				if start > end {
					direction = -1
				}
				for ; start != end; start += direction {
					cave[p.y][start] = "#"
				}
			}
		}
	}
	for row := range cave {
		for col := range cave[row] {
			if cave[row][col] == "" {
				fmt.Print(".")
			} else {
				fmt.Print(cave[row][col])
			}
		}
		fmt.Println()
	}

}
