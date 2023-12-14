package main

import (
	"errors"
)

var (
	FactNegError = errors.New("factorial is not defined for negative numbers")
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, FactNegError
	}

	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res, nil
}
