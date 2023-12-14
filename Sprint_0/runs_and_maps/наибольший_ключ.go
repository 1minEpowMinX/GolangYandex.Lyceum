package main

func FindMaxKey(m map[int]int) int {
	maxKey := 0
	for k := range m {
		if k > maxKey || maxKey == 0 {
			maxKey = k
		}
	}

	return maxKey
}
