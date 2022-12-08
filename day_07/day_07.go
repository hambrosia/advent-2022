package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetTestData(filename string) (data []string, err error) {
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
	return ret, scanner.Err()
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
				i = j
			}
		}
		if strings.HasPrefix(line, "$ cd") && strings.HasSuffix(line, "..") {
			// pop last dir name from path
			path = path[:len(path)-1]
		}
	}

	return folderTree, folderSizes
}

func GetSizeOfFolders(folderTree map[string][]string, folderSizes map[string]int) (p1 int, p2 int) {
	fmt.Println("root size", folderSizes["/"])

	 

	for name := range folderSizes {
		if _, ok := folderTree[name]; !ok {
			fmt.Println("end found", name, folderSizes[name])
			fmt.Println()
			folders := strings.Split(name, ".")
			for i := len(folders); i > 1; i-- {
				folder := strings.Join(folders[:i], ".")
				parent := strings.Join(folders[:i-1], ".")
				fmt.Println("updating parent of", folder)
				fmt.Println("folder has", folderSizes[folder])
				fmt.Println("parent has", folderSizes[parent])
				// add current folder size to parent
				folderSizes[parent] += folderSizes[folder]
				fmt.Println("parent new value", folderSizes[parent])
			}
		} else {
			fmt.Println("skippin")
		}
	}
	fmt.Println("root size", folderSizes["/"])

	// p1 answer
	for _, val := range folderSizes {
		if val < 100000 {
			p1 += val
		}
	}

	fmt.Println("root size", folderSizes["/"])

	// p2 answer
	// foldersBySize := make([]string, 0, len(folderSizes))
	// for name := range folderSizes {
	// 	foldersBySize = append(foldersBySize, name)
	// }

	// sort.SliceStable(foldersBySize, func(i, j int) bool {
	// 	return folderSizes[foldersBySize[i]] < folderSizes[foldersBySize[j]]
	// })

	// for name := range folderSizes {
	// 	fmt.Println(name, folderSizes[name])
	// }

	// for _, name := range foldersBySize {
	// 	originalSize := folderSizes["/"]
	// 	newSize := originalSize - folderSizes[name]
	// 	if 70000000-newSize > 30000000 {
	// 		p2 = folderSizes[name]
	// 		break
	// 	}
	// }
	return p1, p2
}
