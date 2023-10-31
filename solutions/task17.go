package Solution

import (
	"fmt"
	"sort"
)

/*
Задание 17:
Реализовать бинарный поиск встроенными методами языка
*/

func Task17() {

	// Создаем отсортированный массив целых чисел
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	// Ищем индекс числа 12 в отсортированном массиве
	index := sort.SearchInts(arr, 12)

	// Печатаем результат поиска
	if index < len(arr) && arr[index] == 12 {
		fmt.Printf("Число 12 найдено в массиве по индексу %d\n", index)
	} else {
		fmt.Println("Число 12 не найдено в массиве")
	}
}

/*
Тут так же не до конца понял задание, если сортировку нужно написать самостоятельно, тогда мы берем отсортированный
массив и делим его len() на 2 после чего начинаем искать нужное число в середине массива, если искомое число будет больше чем
элемент в серидине, тогда идем искать вправо по массиву, если меньше идем влево.
Пример:

func MySearch(arr []int, index int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == index {
			return mid
		} else if arr[mid] < index {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // Элемент не найден
}

Соответственно при таком коде нужно изменить вывод элемента в Task17() на if index=-1{}else{}
*/
