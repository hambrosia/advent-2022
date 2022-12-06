package day06

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func GetTestData(filename string) (data string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New("Unable to read input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ret string
	for scanner.Scan() {
		ret += scanner.Text()
	}
	return ret, nil
}

func GetStartOfFirstPacket(data string, markerLength int) int {
	for i := markerLength; i < len(data); i++ {
		fmt.Println("i: ", i)
		checked := make(map[string]bool)
		for j := i - markerLength; j <= i; j++ {
			fmt.Println(string(data[j]))
			// check for duplicates
			fmt.Println(checked)
			if !checked[string(data[j])] {
				// no dupe, add to checked
				checked[string(data[j])] = true
			} else {
				fmt.Println("dupe found")
				// dupe found
				break
			}
			if j == i {
				fmt.Println("Success at index: ", i+1)
				return i + 1
			}

		}
		fmt.Println("+++++++++++")
	}

	return -1
}
