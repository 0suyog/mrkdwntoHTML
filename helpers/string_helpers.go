package helpers

import "errors"

func AtPosition(str string, position int) (string, error) {
	runes := []rune(str)
	if position < 0 || position >= len(runes) {
		return "", errors.New("Position out of range")
	}
	return string(runes[position]), nil
}
