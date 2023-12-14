package main

import (
	"fmt"
)

func FindValue(nums []int) int { // *неэффективно
	m := make(map[int]int)

	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = 1
		} else {
			m[n] += 1
		}
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return -1
}

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i] // Используем операцию XOR для обнаружения одиночного числа
	}
	return res
}

func main() {
	fmt.Print(FindValue([]int{2, 2, 1}))
}
