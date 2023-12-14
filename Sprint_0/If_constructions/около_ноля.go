package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	switch {
	case n > 0 && m > 0:
		fmt.Println("Оба числа положительные")
	case n < 0 && m < 0:
		fmt.Println("Оба числа отрицательные")
	case n > 0 && m < 0 || n < 0 && m > 0:
		fmt.Println("Одно число положительное, а другое отрицательное")
	default:
		fmt.Println("Одно из чисел равно нулю")
	}
}
