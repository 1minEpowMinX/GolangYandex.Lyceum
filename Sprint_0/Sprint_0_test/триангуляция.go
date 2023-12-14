package main

import (
	"fmt"
	"strings"
)

func DrawFlower(hB, hS int) {
	if hS > 0 && hB == 0 {
		hB = 1
	}

	for i := 0; i < hB; i++ {
		fmt.Println(strings.Repeat("*", hB-i))
	}

	for i := 1; i < hS; i++ {
		fmt.Println("*")
	}
}

func main() {
	a, b := 0, 0
	fmt.Scanf("%d\n%d", &a, &b)
	DrawFlower(a, b)
}
