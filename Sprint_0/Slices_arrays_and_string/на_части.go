/*
*Сложная задача!

*Напишите функцию CountSubslices(slice []int) int, которая принимает слайс целочисленных элементов,
*находит все подслайсы, сумма элементов которых больше среднего значения суммы элементов всего слайса,
*и возвращает их количество.

*Примечания
*Например, если передать функции CountSubslices(slice []int) слайс [1 2 3] (среднее значение элементов – 2),
*то она должна вернуть 4 (потому что можно взять следующие слайсы: [1 2], [1 2 3], [2 3], [3].
 */

package main

func sumSlice(slice []int) int {
	sum := 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}

	return sum
}

func CountSubslices(slice []int) int {
	length := len(slice)
	AV := sumSlice(slice) / length

	cnt := 0
	for sum, i := 0, 0; i < length; i++ {
		for j := i; j < length; j++ {
			sum += slice[j]
			if sum > AV {
				cnt += 1
			}
		}
		sum = 0
	}

	return cnt
}
