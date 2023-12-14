package main

import (
	"fmt"
	"slices"
	"time"
)

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	res := make([]int, length)

	// Используем res как "отсортированный" массив для записанных значений
	// Проходим по входному массиву и записываем каждое число по его индексу - 1 в res
	for _, v := range nums {
		res[v-1] = v
	}

	// Теперь в res у нас есть массив, в котором пропавшие числа имеют значение 0
	// Проходим по res и "сдвигаем" все значения с нулевым значением в начало массива
	j := 0
	for i := 0; i < length; i++ {
		if res[i] == 0 {
			res[j] = i + 1
			j++
		}
	}

	// Обрезаем res до длины j и возвращаем
	return res[:j]
}

func FindMissingValues(nums []int) []int { // *Мой вариант оказался быстрее предложенного
	found := []int{}

	for n := 1; n < len(nums)+1; n++ {
		if !slices.Contains(nums, n) {
			found = append(found, n)
		}
	}

	return found
}

func main() {
	start := time.Now()
	time.Sleep(time.Second * 2)

	fmt.Print(FindMissingValues([]int{4, 3, 2, 7, 8, 2, 3, 1, 9, 10, 11, 11, 13, 14, 15, 19, 17, 18, 19, 20}))

	elapsed := time.Since(start)
	fmt.Printf("page took for  %s", elapsed)

}
