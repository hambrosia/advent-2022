package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
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
					newFolderName := pathString + "." + fields[1]
					// fmt.Println(newFolderName)
					folderTree[pathString] = append(folderTree[pathString], newFolderName)
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

func GetSizeOfFolders(folderTree map[string][]string, folderSizes map[string]int) (p1 int, p2 int, rootSize int) {
	fmt.Println("root size", folderSizes["/"])

	// fmt.Println(folderTree)

	childlessFolders := make([]string, 0)
	for folderName := range folderSizes {
		if _, found := folderTree[folderName]; !found {
			// childless folder found
			childlessFolders = append(childlessFolders, folderName)
		}
	}

	for i := len(childlessFolders) - 1; i >= 0; {
		folderName := childlessFolders[i]
		// fmt.Println("start", i)
		// fmt.Println("childless folders", childlessFolders)
		fields := strings.Split(folderName, ".")

		parentName := strings.Join(fields[:len(fields)-1], ".")

		folderSizes[parentName] += folderSizes[folderName]

		childlessFolders = append(childlessFolders[:i], childlessFolders[i+1:]...)
		i--

		for j, childName := range folderTree[parentName] {
			if childName == folderName {
				folderTree[parentName] = append(folderTree[parentName][:j], folderTree[parentName][j+1:]...)
			}
			if len(folderTree[parentName]) < 1 {
				// fmt.Println("adding new childless folder", parentName)
				childlessFolders = append(childlessFolders, parentName)
				i++
			}
		}
	}

	// p1 answer
	for _, val := range folderSizes {
		if val < 100000 {
			p1 += val
		}
	}

	// p2 answer
	foldersBySize := make([]string, 0, len(folderSizes))
	for name := range folderSizes {
		foldersBySize = append(foldersBySize, name)
	}

	sort.SliceStable(foldersBySize, func(i, j int) bool {
		return folderSizes[foldersBySize[i]] < folderSizes[foldersBySize[j]]
	})

	maxFileSystem := 70000000
	minFreeSpace := 30000000
	rootSize = folderSizes["/"]

	for _, name := range foldersBySize {
		fmt.Println(name, folderSizes[name])
		if name == "/" {
			continue
		}
		folderSize := folderSizes[name]
		currentFreeSpace := maxFileSystem - rootSize
		newFreeSpace := currentFreeSpace + folderSize
		if newFreeSpace > minFreeSpace {
			p2 = folderSize
			break
		}

	}

	return p1, p2, rootSize
}
