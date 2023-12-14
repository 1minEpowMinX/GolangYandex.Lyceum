package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	switch {
	case n < 0:
		fmt.Println("Некорректный ввод")
	case n < 10:
		fmt.Println("Число меньше 10")
	case n < 100:
		fmt.Println("Число меньше 100")
	case n < 1000:
		fmt.Println("Число меньше 1000")
	default:
		fmt.Println("Число больше или равно 1000")
	}
}
