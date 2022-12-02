package day_01

func getTopCalorieElf(data [][]int) (maxElf int, maxCalCount int) {
	// Given an array of int arrays, return the index of the array with the highest sum

	for i, calorieList := range data {
		tmpCalCount := 0
		for _, calorieCount := range calorieList {
			tmpCalCount += calorieCount
		}
		if tmpCalCount > maxCalCount {
			maxCalCount = tmpCalCount
			maxElf = i
		}
	}
	return maxElf, maxCalCount
}

type elf struct {
	ID       int
	calories int
}

func getTopNElves(data [][]int, n int) (maxElves []elf) {
	for i, calorieList := range data {
		tmpCalCount := 0
		for _, calorieCount := range calorieList {
			tmpCalCount += calorieCount
		}
		if len(maxElves) < 1 {
			maxElves = append(maxElves, elf{i, tmpCalCount})
		} else {
			for j := 0; j < n-1 && j < len(maxElves); j++ {
				if tmpCalCount < maxElves[j].calories {
					continue
				} else {
					//insert here, shift original right, trim slice
					maxElves = append(maxElves[:j+1], maxElves[j:]...)
					maxElves[j] = elf{i, tmpCalCount}
					if len(maxElves) > n {
						maxElves = maxElves[:n]
					}
					break
				}
			}
		}
	}
	return maxElves
}
