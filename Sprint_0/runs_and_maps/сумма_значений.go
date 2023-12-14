package main

func SumOfValuesInMap(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}

	return sum
}
