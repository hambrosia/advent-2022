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

	DebugPrint(debug, "compare before loop %v vs %v", leftSlice, rightSlice)

	for lI, rI := 0, 0; lI < len(leftSlice) || rI < len(leftSlice); lI, rI = lI+1, rI+1 {
		DebugPrint(debug, "peeking to check which slice will run out")
		if rI >= len(rightSlice) && lI < len(leftSlice) {
			DebugPrint(debug, "Right side ran out")
			DebugPrint(debug, "type l %T, type r %T", rightSlice, leftSlice)
			DebugPrint(debug, "left slice %v, right slice %v", leftSlice, rightSlice)
			DebugPrint(debug, "rI %v, lI %v, len r %v, len l %v", rI, lI, len(rightSlice), len(leftSlice))
			// TODO() Start here, why are leftSlice and rightSlice pointing to the outermost function's inputs?
			return false
		} else if rI < len(rightSlice) && lI >= len(leftSlice) {
			DebugPrint(debug, "Left side ran out")
			return true
		}

		DebugPrint(debug, "about to key into slices")
		left, right := leftSlice[lI], rightSlice[rI]
		leftType, rightType := reflect.TypeOf(left).Kind(), reflect.TypeOf(right).Kind()
		DebugPrint(debug, "compare %v vs %v", left, right)

		switch {
		case leftType == reflect.Float64 && rightType == reflect.Float64 && left.(float64) > right.(float64):
			DebugPrint(debug, "both int, left higher than right")
			return false
		case leftType == reflect.Float64 && rightType == reflect.Float64 && left.(float64) < right.(float64):
			DebugPrint(debug, "both int, right higher than left")
			return true
		case leftType == reflect.Slice && rightType == reflect.Slice:
			DebugPrint(debug, "both slice")
			if PacketsInOrder(left.([]interface{}), right.([]interface{}), debug) {
				continue
			} else {
				return false
			}
		case leftType != rightType:
			DebugPrint(debug, "one int, one slice")
			if leftType == reflect.Float64 {
				DebugPrint(debug, "left is num")
				return PacketsInOrder([]interface{}{left}, right.([]interface{}), debug)
			} else {
				DebugPrint(debug, "right is num")
				DebugPrint(debug, "left %v right %v", left, right)
				DebugPrint(debug, "left translate %v, right translate %v", left.([]interface{}), []interface{}{right})
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

		inOrder := PacketsInOrder(left, right, debug)
		DebugPrint(debug, "in order %v ++++++++++", inOrder)
		fmt.Println()

		if inOrder {
			sumRightOrderIndices += index
		}

	}

	return sumRightOrderIndices
}
