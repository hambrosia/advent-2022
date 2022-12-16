package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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

func PacketsInOrder(leftSlice []interface{}, rightSlice []interface{}) (inOrder int) {
	fmt.Printf("compare %v vs %v\n", leftSlice, rightSlice)

	for lI, rI := 0, 0; lI < len(leftSlice) || rI < len(rightSlice); lI, rI = lI+1, rI+1 {

		// peek to see if either packet will run out this cycle
		if rI >= len(rightSlice) && lI < len(leftSlice) {
			return -1
		} else if rI < len(rightSlice) && lI >= len(leftSlice) {
			return 1
		}

		left, right := leftSlice[lI], rightSlice[rI]
		leftType, rightType := reflect.TypeOf(left).Kind(), reflect.TypeOf(right).Kind()
		fmt.Printf("compare %v vs %v\n", left, right)

		switch {
		case leftType == reflect.Float64 && rightType == reflect.Float64 && left.(float64) > right.(float64):
			return -1
		case leftType == reflect.Float64 && rightType == reflect.Float64 && left.(float64) < right.(float64):
			return 1
		case leftType == reflect.Slice && rightType == reflect.Slice:
			packsInOrder := PacketsInOrder(left.([]interface{}), right.([]interface{}))
			if packsInOrder == 0 {
				continue
			} else if packsInOrder == -1 {
				return -1
			} else {
				return 1
			}
		case leftType != rightType:
			if leftType == reflect.Float64 {
				return PacketsInOrder([]interface{}{left}, right.([]interface{}))
			} else {
				return PacketsInOrder(left.([]interface{}), []interface{}{right})
			}
		}
	}
	return 0
}

func SumPacketsInOrder(packets []string) (sumRightOrderIndices int) {
	for i := 0; i < len(packets); i = i + 2 {
		index := (i / 2) + 1
		fmt.Println("Pair", index)

		p1, p2 := packets[i], packets[i+1]

		left, right := []interface{}{}, []interface{}{}
		json.Unmarshal([]byte(p1), &left)
		json.Unmarshal([]byte(p2), &right)

		inOrder := PacketsInOrder(left, right)
		fmt.Println("result", inOrder)
		if inOrder == 1 {
			sumRightOrderIndices += index
			fmt.Println("in order", true)
		} else {
			fmt.Println("in order", false)
		}
		fmt.Println()

	}

	return sumRightOrderIndices
}
