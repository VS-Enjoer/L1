package Solution

import (
	"fmt"
	"sort"
)

/*
Задание 16:
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func Task16() {
	arr := []int{9, 4, 2, 7, 5, 8, 3, 1, 6}
	sort.Ints(arr)

	fmt.Println("Отсортированный массив:", arr)
}

/*
Я не совсем понял задание, если имелось ввиду что нельзя использовать встроенные модули то можно сделать сортировку пузырьком.

func MySort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
*/
