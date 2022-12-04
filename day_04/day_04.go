package day04

func RangeContains(s1 int, e1 int, s2 int, e2 int) bool {
	if s1 <= s2 && e1 >= e2 {
		return true
	}
	if s2 <= s1 && e2 >= e1 {
		return true
	}
	return false
}

func RangeOverlaps(s1 int, e1 int, s2 int, e2 int) bool {
	if s1 <= s2 && (e1 >= e2 || e1 >= s2) {
		return true
	}
	if s2 <= s1 && (e2 >= e1 || e2 >= s1) {
		return true
	}
	return false
}

func FindFullyOverlappingRegions(shifts [][]int) (overlaps int) {
	for _, shift := range shifts {
		s1, e1, s2, e2 := shift[0], shift[1], shift[2], shift[3]
		if RangeContains(s1, e1, s2, e2) {
			overlaps++
		}
	}
	return overlaps
}

func FindPartiallyOverlappingRegions(shifts [][]int) (overlaps int) {
	for _, shift := range shifts {
		s1, e1, s2, e2 := shift[0], shift[1], shift[2], shift[3]
		if RangeOverlaps(s1, e1, s2, e2) {
			overlaps++
		}
	}
	return overlaps
}
