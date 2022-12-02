package day2

func transformThrow(throw string) (string, error) {
	if throw == "A" || throw == "X" {
		return ROCK, nil
	}

	if throw == "B" || throw == "Y" {
		return PAPER, nil
	}

	if throw == "C" || throw == "Z" {
		return SCISSORS, nil
	}

	return NONE, WrongInput
}

func getThrowPoints(playerThrow string) (int, error) {
	if playerThrow == ROCK {
		return 1, nil
	}

	if playerThrow == PAPER {
		return 2, nil
	}

	if playerThrow == SCISSORS {
		return 3, nil
	}

	return -1, WrongInput
}

// elf   A = Rock,  B = Paper,  C = Scissors
// play  X = Rock,  Y = Paper,  Z = Scissors

func getGamePoints(elfThrow, playerThrow string) (int, error) {

	if elfThrow == ROCK {
		if playerThrow == ROCK {
			return 3, nil
		}

		if playerThrow == PAPER {
			return 6, nil
		}

		if playerThrow == SCISSORS {
			return 0, nil
		}
	}

	if elfThrow == PAPER {
		if playerThrow == PAPER {
			return 3, nil
		}

		if playerThrow == SCISSORS {
			return 6, nil
		}

		if playerThrow == ROCK {
			return 0, nil
		}
	}

	if elfThrow == SCISSORS {
		if playerThrow == SCISSORS {
			return 3, nil
		}

		if playerThrow == ROCK {
			return 6, nil
		}

		if playerThrow == PAPER {
			return 0, nil
		}
	}

	return -1, WrongInput
}
