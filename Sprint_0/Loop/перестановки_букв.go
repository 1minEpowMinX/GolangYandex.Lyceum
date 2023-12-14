package main

import "fmt"

func main() {
	var n, total int = 0, 1
	fmt.Scanln(&n)

	for i := 1; i < n; i++ {
		total += total * i
	}
	fmt.Println(total)
}
