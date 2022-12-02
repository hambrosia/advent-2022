package day02

type RPSChoice struct {
	name  string
	beats string
	loses string
	value int
}

func decode(code string) RPSChoice {
	rock := RPSChoice{"rock", "scissors", "paper", 1}
	paper := RPSChoice{"paper", "rock", "scissors", 2}
	scissors := RPSChoice{"scissors", "paper", "rock", 3}

	RPSCodec := map[string]RPSChoice{
		"A":        rock,
		"B":        paper,
		"C":        scissors,
		"X":        rock,
		"Y":        paper,
		"Z":        scissors,
		"rock":     rock,
		"paper":    paper,
		"scissors": scissors,
	}
	return RPSCodec[code]
}

func (choice1 RPSChoice) challenge(choice2 RPSChoice) (choice1Points int, choice2Points int) {
	RPSOutcomeValues := map[string]int{
		"l": 0,
		"d": 3,
		"w": 6,
	}

	if choice1.name == choice2.name {
		return choice1.value + RPSOutcomeValues["d"], choice2.value + RPSOutcomeValues["d"]
	}
	if choice1.beats == choice2.name {
		return choice1.value + RPSOutcomeValues["w"], choice2.value + RPSOutcomeValues["l"]
	} else {
		return choice1Points + RPSOutcomeValues["d"], choice2.value + RPSOutcomeValues["w"]
	}
}

func calculateScore(matches [][]string) (score1 int, score2 int) {
	for _, match := range matches {
		ch1 := decode(match[0])
		ch2 := decode(match[1])
		matchScore1, matchScore2 := ch1.challenge(ch2)
		score1 += matchScore1
		score2 += matchScore2
	}
	return score1, score2
}

func decodeStrategy(code string) string {
	strategyCodec := map[string]string{
		"X": "loses",
		"Y": "draws",
		"Z": "beats",
	}
	return strategyCodec[code]
}

func calculateScoreWithCorrectCodec(matches [][]string) (score1 int, score2 int) {
	for _, match := range matches {
		ch1 := decode(match[0])
		strategy := decodeStrategy(match[1])
		ch2 := RPSChoice{}
		switch strategy {
		case "loses":
			ch2 = decode(ch1.beats)
		case "beats":
			ch2 = decode(ch1.loses)
		default:
			ch2 = decode(ch1.name)
		}
		matchScore1, matchScore2 := ch1.challenge(ch2)
		score1 += matchScore1
		score2 += matchScore2
	}
	return score1, score2
}
