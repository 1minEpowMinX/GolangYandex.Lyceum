package main

// import "slices"

func findMinMaxSlice(a, b []int) ([]int, []int) {
	if len(a) < len(b) {
		return a, b
	} else {
		return b, a
	}
}

func IntersectionOfSlices(slice1, slice2 []int) []int {
	intersecSlice := make([]int, 0, 0)
	min, max := findMinMaxSlice(slice1, slice2)

	for i := 0; i < len(min); i++ {
		for j := 0; j < len(max); j++ {
			if min[i] == max[j] {
				intersecSlice = append(intersecSlice, min[i])
			}
		}
	}

	return intersecSlice
}

// func findMin(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func IntersectionOfSlices(slice1, slice2 []int) []int {
// 	intersecSlice := make([]int, 0, 0)
// 	slices.Sort(slice1)
// 	slices.Sort(slice2)

// 	for i := 0; i < findMin(len(slice1), len(slice2)); i++ {
// 		ind, found := slices.BinarySearch(slice2, slice1[i])
// 		if found {
// 			append(intersecSlice, slice2[ind])
// 		}
// 	}

// 	return intersecSlice
// }
