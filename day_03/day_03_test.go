package day03

import (
	"bufio"
	"os"
	"testing"
)

func TestDay03(t *testing.T) {

	file, err := os.Open("rucksacks.txt")
	if err != nil {
		t.Fatal("Unable to read test input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	testCases := map[string]struct {
		input []string
		want  []int
	}{
		"basic": {
			input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			want: []int{157, 70},
		},
		"longerInput": {
			input: data,
			want:  []int{8243, 2631},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotSum, gotSumGroup := getTotalPriorityOfMisfiledItems(tc.input)
			if gotSum != tc.want[0] {
				t.Errorf("got %v, want %v", gotSum, tc.want[0])
			}
			if gotSumGroup != tc.want[1] {
				t.Errorf("got %v, want %v", gotSumGroup, tc.want[1])
			}
		})
	}
}
