package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 1; i < 11; i++ {
		fmt.Println(n, "*", i, "=", n*i)
	}
}
