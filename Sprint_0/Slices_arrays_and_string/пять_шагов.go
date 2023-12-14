package main

import "slices"

// func ReverseSlice(slice []int) []int {
// 	length := len(slice)
// 	reverseSlice := make([]int, length)

// 	for i := 0; i < length; i++ {
// 		reverseSlice[i] = slice[length-1-i]
// 	}

// 	return reverseSlice
// }

func ReverseSlice(slice []int) []int {
	reverseSlice := make([]int, len(slice))
	copy(reverseSlice, slice)
	slices.Reverse(reverseSlice)

	return reverseSlice
}
