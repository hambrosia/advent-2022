package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name         string
	value        int
	takesCycles  int
	currentCycle int
}

func GetInput(filename string) (res []Instruction) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to read test input.")
		return nil
	}
	defer file.Close()

	nameToCycles := map[string]int{
		"noop": 1,
		"addx": 2,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		val := 0
		if len(line) > 1 {
			val, _ = strconv.Atoi(line[1])
		}
		res = append(res, Instruction{line[0], val, nameToCycles[line[0]], 0})
	}
	return res
}

func CycleIsInteresting(cycleNumber int, target int) bool {
	if cycleNumber == 20 {
		return true
	}
	if cycleNumber > (19+target) && (cycleNumber-20)%target == 0 {
		return true
	}
	return false
}

func SpriteVisible(x int, cycle int) bool {
	beam := (cycle - 1) % 40
	return x-2 < beam && beam < x+2
}

func DisplayOutputToString(displayOutput [240]bool) (ret string) {
	for i, pixel := range displayOutput {
		if i > 39 && i%40 == 0 {
			ret += "\n"
		}
		if pixel {
			ret += "#"
		} else {
			ret += "."
		}

	}
	return ret
}

func doInstruction(instr *Instruction, X *int, cycles *int) (instructionCompleted bool) {
	switch {
	case instr.name == "noop":
		instructionCompleted = true
	case instr.currentCycle != instr.takesCycles:
		instructionCompleted = false
	case instr.name == "addx":
		instructionCompleted = true
		(*X) += instr.value
	}
	return instructionCompleted
}

func Process(data []Instruction) (sigStrength int, displayOutput [240]bool) {
	X := 1
	cycle := 1
	for i := 0; i < len(data); {
		data[i].currentCycle++
		if CycleIsInteresting(cycle, 40) {
			sigStrength += (cycle * X)
		}
		if SpriteVisible(X, cycle) {
			displayOutput[cycle-1] = true
		}
		instructionCompleted := doInstruction(&data[i], &X, &cycle)
		cycle++
		if instructionCompleted {
			i++
		}

	}
	return sigStrength, displayOutput
}
