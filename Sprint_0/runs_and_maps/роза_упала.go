/*
*Палиндром - строка, которая читается одинаково слева направо и справа налево.
*Напишите функцию IsPalindrome(input string) bool, которая принимает строку и проверяет, является ли она палиндромом.

*Примечания
*Например, функция IsPalindrome("А роза упала на лапу Азора") должна вернуть true.

*Вам пригодится пакет strings.
 */

package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(input string) bool {
	inputRune := []rune(strings.ToLower(strings.Replace(input, " ", "", -1)))
	length := len(inputRune)

	for i := 0; i < length/2; i++ {
		if inputRune[i] != inputRune[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(IsPalindrome("А роза упала на лапу Азора"))
}
