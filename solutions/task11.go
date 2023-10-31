package Solution

import "fmt"

/*
Задание 11:
Реализовать пересечение двух неупорядоченных множеств.
*/

func intersection(set1, set2 []int) []int {

	// Создаем пустой срез для хранения результата пересечения
	result := make([]int, 0)

	for _, item1 := range set1 {
		for _, item2 := range set2 {
			if item1 == item2 {
				result = append(result, item1)
				break // Добавить только уникальные элементы пересечения
			}
		}
	}

	return result
}

func Task11() {
	set1 := []int{1, 2, 3, 4, 0}
	set2 := []int{3, 4, 5, 6, 0}

	result := intersection(set1, set2)

	fmt.Println("Результат пересечения:", result)
}
