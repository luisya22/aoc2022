package day6

func markerIsValid(q []string, markerLen int) bool {
	if len(q) != markerLen {
		return false
	}

	groupedMakersMap := map[string]string{}

	for _, c := range q {
		_, alreadyMapped := groupedMakersMap[c]
		if alreadyMapped {
			return false
		} else {
			groupedMakersMap[c] = c
		}
	}

	return true
}
