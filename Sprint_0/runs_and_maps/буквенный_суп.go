/*
*Сложная задача!

*Взяв слово, как пачку пельменей, Гоша высыпает его в кастрюлю и перемешивает.
*Какие слова можно получить таким образом?

*Напишите функцию Permutations(input string) []string, которая принимает строку и выводит все перестановки её
*символов в алфавитном порядке.

*Примечания
*Например, если передать функции Permutations(input string) []string строку "abc" или "cab",
*то она должна вернуть [abc acb bac bca cab cba].

*Для сортировки объектов используйте пакет sort.
 */

package main

import (
	"fmt"
	"sort"
)

func Permutations(input string) []string {
	runes := []rune(input)
	length := len(runes)
	perm := []string{}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j != i {
				runes[i], runes[j] = runes[j], runes[i]
				perm = append(perm, string(runes))
			}
		}
	}

	sort.Strings(perm)
	return perm
}

func main() {
	fmt.Println(Permutations("abc"))
}
