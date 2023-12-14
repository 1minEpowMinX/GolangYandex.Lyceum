package main

import (
	"fmt"
	"slices"
)

// func FindMinMaxInSlice(slice []int) (int, int) {
// 	if slice == nil {
// 		return 0, 0
// 	}

// 	var min, max int = 0, 0
// 	for i := 0; i < len(slice); i++ {
// 		if slice[i] < min || min == 0 {
// 			min = slice[i]
// 		} else if slice[i] > max || max == 0 {
// 			max = slice[i]
// 		}
// 	}

// 	return min, max
// }

func FindMinMaxInSlice(slice []int) (int, int) {
	if len(slice) == 0 {
		return 0, 0
	}

	return slices.Min(slice), slices.Max(slice)
}

func main() {
	fmt.Println(FindMinMaxInSlice([]int{}))
}
