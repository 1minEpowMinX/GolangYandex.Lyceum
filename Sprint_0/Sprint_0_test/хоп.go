package main

import (
	"fmt"
	"math/big"
)

func isPrime(n int64) bool {
	i := big.NewInt(n)
	return i.ProbablyPrime(1)
}

func MashaPlay(place int) {
	for i := 3; i < place; i += 5 {
		if isPrime(int64(i)) {
			fmt.Printf("хоп ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	place := 0
	fmt.Scan(&place)

	MashaPlay(place)
}
