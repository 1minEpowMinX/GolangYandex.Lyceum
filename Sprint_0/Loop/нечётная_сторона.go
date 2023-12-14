package main

import "fmt"

func main() {
	var n, total int
	fmt.Scanln(&n)

	if n < 0 {
		fmt.Println("Некорректный ввод")
		return
	}

	for i := 1; i <= n; i += 2 {
		total += i
	}
	fmt.Println(total)
}
