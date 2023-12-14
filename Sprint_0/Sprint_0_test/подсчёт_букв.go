package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countChar(char, s string) int {
	return strings.Count(s, char)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	chr1, _ := reader.ReadString('\n')
	chr2, _ := reader.ReadString('\n')
	target, _ := reader.ReadString('\n')

	chr1 = strings.TrimSpace(chr1)
	chr2 = strings.TrimSpace(chr2)
	target = strings.TrimSpace(target)

	if a, b := countChar(chr1, target), countChar(chr2, target); a >= b {
		fmt.Printf("%s %d\n", chr1, a)
		fmt.Printf("%s %d", chr2, b)
	} else {
		fmt.Printf("%s %d\n", chr2, b)
		fmt.Printf("%s %d", chr1, a)
	}
}
