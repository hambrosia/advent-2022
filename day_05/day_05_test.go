package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func TestDay05(t *testing.T) {

	file, err := os.Open("large_input.txt")
	if err != nil {
		t.Fatal("Unable to read test input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	crates := [9][]string{}
	operations := []Move{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "move") {
			// fmt.Println("adding operation")
			parts := strings.Fields(line)

			count, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			to, _ := strconv.Atoi(parts[5])
			move := Move{count, from - 1, to - 1}
			operations = append(operations, move)

		} else if len(line) < 1 {
			fmt.Println("skipping empty line")
		} else if strings.HasPrefix(line, " 1") {
			// intended to dynamically get number of stacks here, but ended up just hardcoding it
			fmt.Println("getting number of crate stacks")
			fmt.Printf("number of stacks is: %v \n", line[len(line)-1:])
		} else {
			fmt.Println("adding crates to stack")
			for i := 1; i < len(line); i += 4 {
				character := line[i]
				if unicode.IsUpper(rune(character)) {
					stackNumber := i / 4
					fmt.Printf("found crate in stack %v\n", stackNumber)
					// prepend found character to stack
					crates[stackNumber] = append([]string{string(character)}, crates[stackNumber]...)
				}

			}

		}

	}
	// fmt.Println(crates)

	// large input p1 want CVCWCRTVQ
	// small input p1 want CMZ

	// large input p2 want CNSCZWLVT
	// small input p2 want MCD

	wantP1 := "CVCWCRTVQ"
	gotP1 := ArrangeCratesAndGetTopmost(crates, operations, 9000)
	if gotP1 != wantP1 {
		t.Errorf("p1: got %v, want %v", gotP1, wantP1)
	}

	gotP2 := ArrangeCratesAndGetTopmost(crates, operations, 9001)
	wantP2 := "CNSCZWLVT"
	if gotP1 != wantP1 {
		t.Errorf("p2: got %v, want %v", gotP2, wantP2)
	}
}
