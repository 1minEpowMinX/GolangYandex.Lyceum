package main

import "errors"

var (
	OutOfRangeError = errors.New("position out of range")
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	runes := []rune(str)

	if position < 0 || position > len(runes)-1 {
		return 0, OutOfRangeError
	}

	return runes[position], nil
}
