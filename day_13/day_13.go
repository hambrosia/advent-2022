package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	. "github.com/hambrosia/advent-2022/helpers"
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

func PacketsInOrder(leftSlice []interface{}, rightSlice []interface{}, debug bool) (inOrder bool) {

	for lI, rI := 0, 0; lI < len(leftSlice) && rI < len(leftSlice); lI, rI = lI+1, rI+1 {
		left, right := leftSlice[lI], rightSlice[rI]
		leftType, rightType := reflect.TypeOf(left).Kind(), reflect.TypeOf(right).Kind()

		switch {
		case leftType == reflect.Float64 && rightType == reflect.Float64 && left.(float64) > right.(float64):
			return false
		case leftType == reflect.Slice && rightType == reflect.Slice:
			return PacketsInOrder(left.([]interface{}), right.([]interface{}), debug)
		case leftType != rightType:
			if leftType == reflect.Float64 {
				return PacketsInOrder([]interface{}{left}, right.([]interface{}), debug)
			} else {
				DebugPrint(debug, "left %v %T", left, left)
				DebugPrint(debug, "right %v %T", right, right)
				return PacketsInOrder(left.([]interface{}), []interface{}{right}, debug)
			}
		}

		// TODO() check indexes and lengths
	}

	return true
}

func SumPacketsInOrder(packets []string, debug bool) (sumRightOrderIndices int) {
	for i := 0; i < len(packets); i += 2 {
		p1 := packets[i]
		p2 := packets[i+1]
		index := ((i + 2) / 3) + 1

		left, right := []interface{}{}, []interface{}{}
		json.Unmarshal([]byte(p1), &left)
		json.Unmarshal([]byte(p2), &right)

		DebugPrint(debug, "pair %v", index)
		inOrder := PacketsInOrder(left, right, debug)
		DebugPrint(debug, "in order %v", inOrder)

		if inOrder {
			sumRightOrderIndices += index
		}

	}

	return sumRightOrderIndices
}