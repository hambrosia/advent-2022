package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(filename string) (heightmap [][]int, start Point, end Point) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening map file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		bytes := scanner.Bytes()
		line := []int{}
		for j, byte := range bytes {
			val := 0
			switch byte {
			case 83:
				start = Point{j, i}
			case 69:
				end = Point{j, i}
				val = 25
			default:
				val = int(byte) - 97
			}
			line = append(line, val)
		}
		heightmap = append(heightmap, line)
		i++
	}
	return heightmap, start, end
}

type Point struct {
	x int
	y int
}

func (p Point) GetNeighbors(heightmap [][]int) (neighbors []Point) {
	modifiers := [][]int{
		{0, -1}, // 12:00
		{1, 0},  // 3:00
		{0, 1},  // 6:00
		{-1, 0}, // 9:00
	}
	for _, modifier := range modifiers {
		x, y := p.x-modifier[0], p.y-modifier[1]
		if x < 0 || x >= len(heightmap[0]) {
			continue
		}
		if y < 0 || y >= len(heightmap) {
			continue
		}
		neighbors = append(neighbors, Point{x, y})
	}
	return neighbors
}

func (p Point) FilterNeighborsByHeight(heightmap [][]int, neighbors []Point, graph map[Point][]Point) (reachableNeighbors []Point) {
	pHeight := heightmap[p.y][p.x]
	for _, n := range neighbors {
		nHeight := heightmap[n.y][n.x]
		if nHeight <= pHeight+1 {
			reachableNeighbors = append(reachableNeighbors, n)
		}
	}
	return reachableNeighbors
}

func (p Point) FilterNeighborsByVisited(neighbors []Point, graph map[Point][]Point) (unvisitedNeighbors []Point) {
	for _, n := range neighbors {
		if _, found := graph[n]; found {
			continue
		}
		unvisitedNeighbors = append(unvisitedNeighbors, n)

	}
	return unvisitedNeighbors
}

func PrintPath(heightmap [][]int, path []Point) {
	pathMap := map[Point]string{}
	for _, p := range path {
		pathMap[p] = "X"
	}
	for i := range heightmap {
		for j := range heightmap[i] {
			if _, found := pathMap[Point{j, i}]; found {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")

			}
		}
		fmt.Print("\n")
	}
}

func FindShortestPath(heightmap [][]int, start Point, end Point) (numSteps int, path []Point) {
	// bfs requires input to be a graph, we can try making it as we go along
	// the FilterNeighborsByHeight function can be used as the adjacency list for a node

	// use slice as FIFO queue
	queue := []Point{}
	queue = append(queue, start)
	// graph is an adjacency list of points to reachable neighbor points

	graph := map[Point][]Point{} // just used to keep track of visited, but could be useful in p2
	parents := map[Point]Point{}

	endFound := false
	count := 0
	for len(queue) > 0 {
		// pop a node from the front of the queue
		node := queue[0]
		queue = queue[1:]
		count++
		// fmt.Println("cycle", count)
		// fmt.Println("node", node)
		// fmt.Println("queue start", queue)
		// fmt.Println("graph start", graph)

		// if _, visited := graph[node]; visited {
		// 	continue
		// }

		if node == end {
			fmt.Println("FOUND END!", end)
			graph[node] = []Point{}
			endFound = true
			break
		}

		neighbors := node.GetNeighbors(heightmap)
		neighbors = node.FilterNeighborsByVisited(neighbors, graph)
		neighbors = node.FilterNeighborsByHeight(heightmap, neighbors, graph)

		// fmt.Println("neighbors", neighbors)

		// for each neighbor, if the neighbor is reachable and unvisited add the neighbor to the node's adjacency list
		// add current node as parent of neighbor
		for _, neighbor := range neighbors {
			parents[neighbor] = node
			queue = append(queue, neighbor)
			graph[node] = append(graph[node], neighbor)
		}

		// fmt.Println("queue end", queue)
		// fmt.Println()

	}

	if endFound {
		path = []Point{start, end}
		for node, ok := parents[end]; ok && node != start; node, ok = parents[node] {
			path = append([]Point{node}, path...) // prepend point to path
			PrintPath(heightmap, path)
			fmt.Println()
		}
		numSteps = len(path) - 1
	}

	return numSteps, path
}
