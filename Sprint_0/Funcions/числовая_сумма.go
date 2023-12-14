package main

func SumDigitsRecursive(n int) int {
	total := 0
	if n < 1 {
		return total
	}

	total += n % 10
	return total + SumDigitsRecursive(n/10)
}
