package day3

import "errors"

var InvalidCharacter = errors.New("invalid character")

func getPriority(char rune) (int, error) {

	code := int(char)

	if code >= 97 && code <= 122 {
		return code - 96, nil
	} else if code >= 65 && code <= 90 {
		return code - 38, nil
	}

	return -1, InvalidCharacter
}
