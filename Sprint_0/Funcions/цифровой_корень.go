/*
*Цифровой корень числа - это сумма всех его цифр до тех пор, пока не останется одна цифра.
*Например, цифровой корень числа 12345 равен 1 + 2 + 3 + 4 + 5 = 15, а затем 1 + 5 = 6.

*Напишите функцию CalculateDigitalRoot(n int) int, которая будет принимать на вход положительное целое число n
*и возвращать его цифровой корень.
 */

package main

func CalculateDigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	total := 0
	for n > 0 {
		total += n % 10
		n /= 10
	}

	return CalculateDigitalRoot(total)
}
