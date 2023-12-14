package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	switch {
	case n%2 == 0 && n > 0:
		fmt.Println("Число положительное и четное")
	case n%2 == 0 && n < 0:
		fmt.Println("Число отрицательное и четное")
	case n%2 != 0 && n > 0:
		fmt.Println("Число положительное и нечетное")
	case n%2 != 0 && n < 0:
		fmt.Println("Число отрицательное и нечетное")
	}
}
