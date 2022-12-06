package day05

type Move struct {
	count int
	from  int
	to    int
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func ArrangeCratesAndGetTopmost(crates [9][]string, moves []Move, crateMoverVersion int) (top string) {
	for _, move := range moves {
		sourceCrate := crates[move.from]
		moveCount := Min(move.count, len(sourceCrate))
		toMove := sourceCrate[len(sourceCrate)-moveCount:]
		toMoveRev := []string{}
		switch crateMoverVersion {
		case 9000:
			for i := len(toMove) - 1; i >= 0; i-- {
				toMoveRev = append(toMoveRev, toMove[i])
			}
			crates[move.to] = append(crates[move.to], toMoveRev...)
		case 9001:
			crates[move.to] = append(crates[move.to], toMoveRev...)

		}
		crates[move.from] = crates[move.from][:len(sourceCrate)-moveCount]
	}

	ret := ""
	for _, crate := range crates {
		if len(crate) > 0 {
			ret += crate[len(crate)-1]
		}
	}
	return ret
}
