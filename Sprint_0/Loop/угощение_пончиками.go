package main

import "fmt"

func main() {
	var n, total int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		if i%3 != 0 && i%5 != 0 {
			total += i
		}
	}

	fmt.Println(total)
}
