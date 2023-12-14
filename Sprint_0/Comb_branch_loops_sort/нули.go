package main

import (
	"fmt"
	"time"
)

func moveZeroes(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j { // Проверяем, что i и j не равны
				nums[j] = nums[i]
				nums[i] = 0
			}
			j++
		}
	}

	return nums
}

func MoveZeroes(nums []int) []int {
	length := len(nums)
	for i := 0; i < length/2; i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}

	return nums
}

func main() {
	start := time.Now()
	time.Sleep(time.Second * 2)

	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))

	elapsed := time.Since(start)
	fmt.Printf("page took for  %s", elapsed)
}
