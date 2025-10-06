package helpers

import (
	"errors"
	"fmt"
)

func AtPosition(str string, position int) (string, error) {
	runes := []rune(str)
	if position < 0 || position >= len(runes) {
		return "", errors.New("Position out of range")
	}
	return string(runes[position]), nil
}

func Printf(format string, a ...any) {
	fmt.Println("____________________")
	fmt.Printf(format, a...)
	fmt.Println("____________________")
}
