package Solution

import "fmt"

/*
Задание 23:
Удалить i-ый элемент из слайса
*/

func removeElement(s []int, index int) []int {
	if index < 0 || index >= len(s) { // Возвращаю исходный слайс, если индекс находится за его пределами.
		fmt.Println("Индекс находится за пределами слайса")
		return s
	} else {
		return append(s[:index], s[index+1:]...) // Собираю слайс до и после удаляемого элемента и объединяем их.
	}
}

func Task23() {
	slice := []int{1, 2, 3, 4, 5}
	index := 6

	result := removeElement(slice, index)

	fmt.Println("Слайс после удаления элемента: ", result)
}
