package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(filename string) (res []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read test input from", filename)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		res = append(res, line)
	}
	return res
}

func PackToStack(packet string) (values []int, separators []string) {
	separator := map[string]bool{
		"[": true,
		"]": true,
		",": false,
	}
	for i := 0; i < len(packet); i++ {
		b := packet[i]
		switch {
		case separator[string(b)]:
			separators = append(separators, string(b))
		case string(b) == "1" && string(packet[i+1]) == "0":
			// 10
			values = append(values, 10)
			i++
		case 48 <= b && b <= 57:
			values = append(values, int(b-48))
		}
	}
	return values, separators
}

func ComparePackets(packets []string) (sumRightOrderIndices int) {
	// for each packet pair (list of list or int)
	// left side is p1, right side is p2
	// condition 1
	// if both values are integers and left is not less than right, fail

	// condition 2
	// if both values are lists, compare the first value of each value in each list and continue
	// if left list runs out of items first the inputs are in the right order
	// if right list runs out of items first fail
	// if lists have same length and are identical, move to the next part of the input

	// condition 3
	// if exactly one value is an integer, convert the integer to a list with only that integer as its value, then compare again

	// TLDR
	// as soon as a comparison shows the left side is smaller, the packets are in the right order and you can continue
	// as soon as a comparison shows the right side is smaller, the packers are in the WRONG order, mark and continue

	// pairs are index from 1, p1 has index 1, p2 index 2 etc.

	// return sum of indices of packets in the right order

	for i := 0; i < len(packets); i += 2 {
		p1 := packets[i]
		p2 := packets[i+1]
		index := ((i + 2) / 3) + 1
		fmt.Println("index", index)
		fmt.Println("pair", p1, p2)

		// p2Values := []int{}
		// p2Separators := []string{}

		p1Values, p1Separators := PackToStack(p1)
		fmt.Println("p1 separators", p1Separators)
		fmt.Println("p1 values", p1Values)
		p2Values, p2Separators := PackToStack(p2)
		fmt.Println("p2 separators", p2Separators)
		fmt.Println("p2 values", p2Values)

	}

	return sumRightOrderIndices
}
