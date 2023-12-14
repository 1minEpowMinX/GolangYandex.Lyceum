package main

func SumOfSlice(slice []int) int {
	sum := 0

	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}

	return sum
}
