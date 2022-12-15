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

	DebugPrint(debug, "leftSlice %v", leftSlice)
	DebugPrint(debug, "rightSlice %v", rightSlice)

	DebugPrint(debug, "l: %v, r %v", leftSlice, rightSlice)
	for lI, rI := 0, 0; lI < len(leftSlice) && rI < len(leftSlice); lI, rI = lI+1, rI+1 {
		DebugPrint(debug, "peeking to check which slice will run out")
		if rI >= len(rightSlice) && lI < len(leftSlice) {
			return false
		} else if rI < len(rightSlice) && lI >= len(leftSlice) {
			return true
		}
		DebugPrint(debug, "about to key into slices")
		left, right := leftSlice[lI], rightSlice[rI]
		leftType, rightType := reflect.TypeOf(left).Kind(), reflect.TypeOf(right).Kind()
		DebugPrint(debug, "left %v", left)
		DebugPrint(debug, "right %v", right)
		switch {
		case leftType == reflect.Float64 && rightType == reflect.Float64 && left.(float64) > right.(float64):
			DebugPrint(debug, "both int")
			return false
		case leftType == reflect.Slice && rightType == reflect.Slice:
			DebugPrint(debug, "both slice")
			return PacketsInOrder(left.([]interface{}), right.([]interface{}), debug)
		case leftType != rightType:
			DebugPrint(debug, "one int, one slice")
			if leftType == reflect.Float64 {
				return PacketsInOrder([]interface{}{left}, right.([]interface{}), debug)
			} else {
				return PacketsInOrder(left.([]interface{}), []interface{}{right}, debug)
			}
		default:
			DebugPrint(debug, "didn't hit a condition")
		}

	}

	return true
}

func SumPacketsInOrder(packets []string, debug bool) (sumRightOrderIndices int) {
	for i := 0; i < len(packets); i = i + 2 {
		index := (i / 2) + 1
		DebugPrint(debug, "pair %v ++++++++++", index)
		p1 := packets[i]
		p2 := packets[i+1]

		left, right := []interface{}{}, []interface{}{}
		json.Unmarshal([]byte(p1), &left)
		json.Unmarshal([]byte(p2), &right)

		// DebugPrint(debug, "p1 %v", p1)
		// DebugPrint(debug, "p2 %v", p2)
		DebugPrint(debug, "left %v", left)
		DebugPrint(debug, "right %v", right)
		inOrder := PacketsInOrder(left, right, debug)
		DebugPrint(debug, "in order %v ++++++++++", inOrder)

		if inOrder {
			sumRightOrderIndices += index
		}

	}

	return sumRightOrderIndices
}
