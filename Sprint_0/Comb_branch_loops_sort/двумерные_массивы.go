package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	res := [][]int{}

	if len(original) == m*n {
		for i := 0; i < m; i++ {
			res = append(res, original[i*n:(i+1)*n])
		}
	}

	return res
}

func Convert2D(nums []int, m, n int) [][]int { // *неэффективно
	converted := make([][]int, m)

	for i := 0; i < m; i++ {
		innerSlice := make([]int, n)
		for j := 0; j < n; j++ {
			innerSlice[j] = nums[i*n+j]
		}
		converted[i] = innerSlice
	}

	return converted
}

func main() {
	fmt.Print(Convert2D([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2))
}
