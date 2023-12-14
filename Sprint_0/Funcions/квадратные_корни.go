package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)

	D := math.Sqrt(math.Pow(b, 2) - 4*a*c)

	if D > 0 {
		x1 := (-b + D) / (2 * a)
		x2 := (-b - D) / (2 * a)
		fmt.Println(math.Min(x1, x2), math.Max(x1, x2))
	} else if D == 0 {
		x1 := -b / (2 * a)
		fmt.Println(x1)
	} else {
		fmt.Println(0, 0)
	}
}
