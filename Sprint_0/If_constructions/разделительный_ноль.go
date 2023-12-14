package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	if n > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число отрицательное")
	}
}
