package main

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func PrintNumbersAscending(n int) {
	if n < 1 {
		return
	}

	PrintNumbersAscending(n - 1)
	fmt.Printf("%d ", n)
}
