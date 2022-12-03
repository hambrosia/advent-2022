package day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCalculateScore(t *testing.T) {

	file, err := os.Open("matches.txt")
	if err != nil {
		t.Fatal("Unable to read test input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([][]string, 0)
	index := 0
	for scanner.Scan() {
		choices := strings.Fields(scanner.Text())
		data = append(data, choices)
		index++
	}

	testCases := map[string]struct {
		input [][]string
		want  []int
	}{
		"example": {
			input: [][]string{{"A", "Y"}, {"B", "X"}, {"C", "Z"}},
			want:  []int{14, 15},
		},
		"largeExample": {
			input: data,
			want:  []int{9188, 14163},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got1, got2 := calculateScore(tc.input)
			if got1 != tc.want[0] {
				t.Errorf("Got %v, want %v", got1, tc.want[0])
			}
			if got2 != tc.want[1] {
				t.Errorf("Got %v, want %v", got2, tc.want[1])
			}
			fmt.Println(name, got1, got2)
		})
	}

	correctCodecTestCases := map[string]struct {
		input [][]string
		want  []int
	}{
		"example": {
			input: [][]string{{"A", "Y"}, {"B", "X"}, {"C", "Z"}},
			want:  []int{12, 12},
		},
		"largeExample": {
			input: data,
			want:  []int{11538, 12091},
		},
	}

	for name, tc := range correctCodecTestCases {
		t.Run(name, func(t *testing.T) {
			got1, got2 := calculateScoreWithCorrectCodec(tc.input)
			if got1 != tc.want[0] {
				t.Errorf("Got %v, want %v", got1, tc.want[0])
			}
			if got2 != tc.want[1] {
				t.Errorf("Got %v, want %v", got2, tc.want[1])
			}
			fmt.Println(name, got1, got2)
		})
	}

}
