/*
*Гоша решил провести эксперимент, чтобы понять, как быстро суммируется ряд из дробей.
*Для этого он хочет написать программу, которая по заданному числу n будет вычислять значение
*следующей последовательности: 1 + 1/2 + 1/3 + ... + 1/n.

*Напишите функцию CalculateSeriesSum(n int) float64, которая получает на вход число n и
*возвращает сумму указанной последовательности.
 */

package main

func CalculateSeriesSum(n int) float64 {
	total := 0.0
	if n < 1 {
		return total
	}

	total += 1.0 / float64(n)
	return total + CalculateSeriesSum(n-1)
}
