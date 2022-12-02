package day_01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestGetTopCalorieElf(t *testing.T) {
	file, err := os.Open("calories.txt")
	if err != nil {
		t.Fatal("Unable to read test input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([][]int, 1)
	index := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			data = append(data, []int{})
			index++
		} else {
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				t.Fatal("Unable to read test input")
			}
			data[index] = append(data[index], val)
		}
	}
	gotElf, gotCalories := getTopCalorieElf(data)

	wantElf := 179
	wantCalories := 66719

	if gotElf != wantElf {
		t.Fatal("Got", gotElf, "want", wantElf)
	}
	if gotCalories != wantCalories {
		t.Fatal("Got", gotCalories, "want", wantCalories)
	}

	gotTopThreeElves := getTopNElves(data, 3)

	gotSum := 0
	for _, elf := range gotTopThreeElves {
		gotSum += elf.calories
	}
	wantSum := 198551
	if gotSum != wantSum {
		t.Fatal("Got", gotSum, "want", wantSum)
	}
	fmt.Println(gotTopThreeElves)
}
