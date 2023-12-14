package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scanln(&x, &y, &z)

	if x < 0 || y < 0 || z < 0 {
		fmt.Println("Некорректный ввод")
		return
	}
	if x == y && y == z {
		fmt.Println("Все числа равны")
	} else if x == y || y == z || x == z {
		fmt.Println("Два числа равны")
	} else {
		fmt.Println("Все числа разные")
	}
}
