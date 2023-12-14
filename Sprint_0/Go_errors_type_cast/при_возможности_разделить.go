package main

import "errors"

var (
	DivByZero = errors.New("division by zero is not allowed")
)

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, DivByZero
	}

	return float64(a) / float64(b), nil
}
