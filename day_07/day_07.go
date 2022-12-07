package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	// "sort"
	"strconv"
	"strings"
)

func GetTestData(filename string) (data []string, err error) {
	fmt.Println("In get test data")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return []string{}, errors.New("Unable to read input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret, nil
}

func ParseTestData(data []string) (map[string][]string, map[string]int) {
	// folder name to next level subfolders (adjacent nodes)
	folderTree := make(map[string][]string)
	// size / value of nodes
	folderSizes := make(map[string]int)

	path := make([]string, 0)
out:
	for i, line := range data {
		if strings.HasPrefix(line, "$ cd") && !strings.HasSuffix(line, "..") {
			fields := strings.Fields(line)
			folderName := fields[2]
			path = append(path, folderName)
			pathString := strings.Join(path, ".")
			fmt.Println("path string", pathString)

			for j := i + 2; !strings.HasPrefix(data[j], "$") && j < len(data); j++ {
				fields := strings.Fields(data[j])

				if _, err := strconv.Atoi(string(data[j][0])); err == nil {

					// file found, add to current folder size
					val, _ := strconv.Atoi(fields[0])
					folderSizes[pathString] += val

				}

				if strings.HasPrefix(data[j], "dir") {
					// folder found add to folder tree
					folderTree[pathString] = append(folderTree[pathString], fields[1])
				}
				if j == len(data)-1 {
					break out
				}
			}
		}
		if strings.HasPrefix(line, "$ cd") && strings.HasSuffix(line, "..") {
			// pop last dir name from path
			path = path[:len(path)-1]
			fmt.Println("trimmed path to", path)
		}
	}
	// fmt.Println(folderSizes)
	// fmt.Println(folderTree)

	return folderTree, folderSizes
}

func GetSizeOfFolders(folderTree map[string][]string, folderSizes map[string]int) int {
	queue := make([]string, 0)
	queue = append(queue, "/") // add start to queue

	parents := make(map[string]string)
	parents["/"] = ""

	// for len(queue) > 0 {
	// 	node := queue[0]
	// 	queue = queue[1:]
	// 	for _, v := range folderTree[node] {
	// 		if _, ok := parents[v]; !ok {
	// 			// add enclosing folder as parent
	// 			parents[v] = node

	// 			// add current folder to queue
	// 			queue = append(queue, v)
	// 		}
	// 		if _, ok := folderTree[v]; !ok {
	// 			fmt.Println("end found, update all parents of", v)
	// 			fmt.Println("end folder has size:", folderSizes[v])
	// 			prev := v
	// 			for node, ok := parents[v]; ok && node != ""; node, ok = parents[node] {
	// 				fmt.Println()
	// 				fmt.Println("updating folder", node)
	// 				fmt.Println("current folder size", folderSizes[node])
	// 				fmt.Println("with child folder", v)
	// 				fmt.Println("child folder size", folderSizes[prev])
	// 				folderSizes[node] += folderSizes[prev]
	// 				fmt.Println("new current folder size", folderSizes[node])
	// 				// fmt.Println("press enter to continue")
	// 				// input := bufio.NewScanner(os.Stdin)
	// 				// input.Scan()
	// 				prev = node
	// 			}
	// 		}
	// 	}
	// }
	for folderString := range folderSizes {
		fmt.Println("working on folderString", folderString)
		folders := strings.Split(folderString, ".")
		for i := len(folders); i > 1; i-- {
			folder := strings.Join(folders[:i], ".")
			size := folderSizes[folder]
			fmt.Println("folder", folder)
			fmt.Println("size", size)
			parent := strings.Join(folders[:i-1], ".")

			// add current folder size to parent
			folderSizes[parent] += folderSizes[folder]
			fmt.Println()
		}
	}
	// fmt.Println(folderSizes)

	sum := 0
	for _, val := range folderSizes {
		if val < 100000 {
			sum += val
		}
	}
	// foldersBySize := make([]string, 0, len(folderSizes))

	// fmt.Println(foldersBySize)
	// for name := range folderSizes {
	// 	foldersBySize = append(foldersBySize, name)
	// }
	// fmt.Println("\n")
	// fmt.Println(folderSizes)

	// sort.SliceStable(foldersBySize, func(i, j int) bool {
	// 	return folderSizes[foldersBySize[i]] > folderSizes[foldersBySize[j]]
	// })

	// for _, name := range foldersBySize {
	// 	originalSize := folderSizes["/"]
	// 	fmt.Println()
	// 	fmt.Println("size of folder", folderSizes[name])
	// 	fmt.Println("new size", originalSize-folderSizes[name])
	// }
	// fmt.Println(folderSizes)
	return sum

}

func main() {
	testInput, _ := GetTestData("large_input.txt")
	fmt.Println(testInput, testInput)
	tree, sizes := ParseTestData(testInput)
	res := GetSizeOfFolders(tree, sizes)

	fmt.Println("Answer")
	fmt.Println(res)
	// not 1609610
}
