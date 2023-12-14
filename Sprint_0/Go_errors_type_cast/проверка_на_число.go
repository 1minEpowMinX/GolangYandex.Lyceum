package main

import (
	"errors"
	"strconv"
)

var (
	notInt = errors.New("invalid input, please provide two integers")
)

func SumTwoIntegers(a, b string) (int, error) {
	intA, errA := strconv.Atoi(a)
	intB, errB := strconv.Atoi(b)

	if errA != nil || errB != nil {
		return 0, notInt
	}

	return intA + intB, nil
}
