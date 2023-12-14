/*
*Задача для тех кто хорошо справляется с курсом.

*После урока про сортировки Маша решила применить полученные знания с пользой.

*У неё есть тетрадка, куда она записывает всех, кому за что-нибудь благодарна.
*Список может быть очень длинным, но Маша его упорядочила по алфавиту (это называется - лексикографический порядок).

*Помогите ей организовать быстрый поиск по этому списку.

*Формат ввода
*На первой строке передается количество людей в списке - натуральное число.
*Затем на следующих строках идет список людей. А после списка - перечень префиксов всех тех, кого надо найти.
*Перечень префиксов может быть произвольной длины. Пустая строка вместо префикса значит что префиксы закончились.

*Формат вывода
*Для каждого префикса необходимо вывести первую фамилию с таким префиксом, или фразу Не найдено.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func get_search(search string, slice []string) string {
	for _, people := range slice {
		if strings.Contains(people, search) {
			return people
		}
	}

	return "Не найдено"
}

func main() {
	var n int
	fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)

	slice := []string{}
	for i := 0; i < n; i++ {
		tmp, _ := reader.ReadString('\n')
		tmp = strings.TrimSpace(tmp)

		slice = append(slice, tmp)
	}

	for {
		tmp, _ := reader.ReadString('\n')
		tmp = strings.TrimSpace(tmp)
		if tmp == "" {
			break
		}
		fmt.Println(get_search(tmp, slice))
	}
}
