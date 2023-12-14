package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	if n > m {
		fmt.Println("Первое число больше второго")
	} else if m > n {
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}
}
