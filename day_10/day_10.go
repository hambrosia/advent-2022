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

func SpriteVisible(x int, cycles int) bool {
	return x-2 < cycles && cycles < x+2
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
	cycles := 1
	for i := 0; i < len(data); {
		data[i].currentCycle++

		if CycleIsInteresting(cycles, 40) {
			fmt.Println("Instruction", i)
			fmt.Println("sigStrength", sigStrength)
			fmt.Println("cycles", cycles, "register", X)
			sigStrength += (cycles * X)
			fmt.Println("cycles", cycles, "register", X)
			fmt.Println("sigStrength", sigStrength)
			fmt.Println()
		}
		if SpriteVisible(X, cycles) {
			displayOutput[X] = true
		}
		instructionCompleted := doInstruction(&data[i], &X, &cycles)
		cycles++
		if instructionCompleted {
			i++
		}

	}
	return sigStrength, displayOutput
}

func main() {
	instructions := GetInput("large_input.txt")
	sigStrength, _ := Process(instructions)
	fmt.Println(sigStrength)

}
