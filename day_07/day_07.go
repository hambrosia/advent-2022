package day07

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

func parseTestData(data []string) (map[string][]string, map[string]int) {
	// folder name to next level subfolders (adjacent nodes)
	folderTree := make(map[string][]string)
	// size / value of nodes
	folderSizes := make(map[string]int)
out:
	for i, line := range data {
		if strings.HasPrefix(line, "$ cd") && !strings.HasSuffix(line, "..") {
			fields := strings.Fields(line)
			folderName := fields[2]
			folderSizes[folderName] = 0
			for j := i + 2; !strings.HasPrefix(data[j], "$") && j < len(data); j++ {
				fields := strings.Fields(data[j])

				if _, err := strconv.Atoi(string(data[j][0])); err == nil {

					// file found, add to current folder size
					val, _ := strconv.Atoi(fields[0])
					folderSizes[folderName] += val

				}

				if strings.HasPrefix(data[j], "dir") {
					// folder found add to folder tree
					folderTree[folderName] = append(folderTree[folderName], fields[1])
				}
				if j == len(data)-1 {
					break out
				}
			}

		}
	}
	fmt.Println(folderSizes)
	fmt.Println(folderTree)

	return folderTree, folderSizes
}

func getSizeOfFolders(folderTree map[string][]string, folderSizes map[string]int) int {
	queue := make([]string, 0)
	queue = append(queue, "/") // add start to queue

	parents := make(map[string]string)
	parents["/"] = ""

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, v := range folderTree[node] {
			if _, ok := parents[v]; !ok {
				// add enclosing folder as parent
				parents[v] = node

				// add current folder to queue
				queue = append(queue, v)
			}
			if _, ok := folderTree[v]; !ok {
				fmt.Println("end found, update all parents of", v)
				prev := v
				for node, ok := parents[v]; ok && node != ""; node, ok = parents[node] {
					fmt.Println("updating node", node)
					fmt.Println("with node", v)
					folderSizes[node] += folderSizes[prev]
					prev = node
				}
			}
		}
	}
	sum := 0
	for key, val := range folderSizes {
		if val < 100000 {
			fmt.Println(key, val)
			sum += val
		}
	}
	fmt.Println(folderSizes)
	return sum

}
