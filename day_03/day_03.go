package day03

import (
	"strings"
	"unicode"
)

func getPriorityOfCharacter(character rune) int {
	if unicode.IsUpper(character) {
		return int(character) - 38
	}
	return int(character) - 96
}

func getTotalPriorityOfMisfiledItems(data []string) (sumPriority int, sumGroupPriority int) {

	for i, rucksack := range data {
		mid := (len(rucksack) / 2)
		for _, character := range rucksack[:mid] {
			if strings.Contains(rucksack[mid:], string(character)) {
				priority := getPriorityOfCharacter(character)
				sumPriority += priority
				break
			}
		}
		if ((i + 1) % 3) == 0 {
			// check for common item in group
			for _, character := range rucksack {
				if strings.Contains(data[i-1], string(character)) && strings.Contains(data[i-2], string(character)) {
					priority := getPriorityOfCharacter(character)
					sumGroupPriority += priority
					break
				}
			}
		}
	}
	return sumPriority, sumGroupPriority
}
