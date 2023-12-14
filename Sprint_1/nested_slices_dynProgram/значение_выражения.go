package main

import "fmt"

func calculateExpressionValue(nums []int) int {
	first := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		first[i] = max(first[i+1], nums[i])
	}

	diff := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		diff[i] = max(diff[i+1], first[i]-nums[i])
	}

	return diff[0]
}

func MaxExpressionValue(nums []int) int {
	n := len(nums) / 2
	left, right := nums[:n], nums[n:]

	return calculateExpressionValue(left) + calculateExpressionValue(right)
}

func main() {
	nums := []int{3, 9, 10, 1, 30, 40} // * функция должна вернуть значение 46 (поскольку 40 – 1 + 10 – 3 - максимально)
	fmt.Print(MaxExpressionValue(nums))
}
