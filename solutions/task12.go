package Solution

import "fmt"

/*
Задание 12:
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func createSet(arr []string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, item := range arr {
		set[item] = struct{}{} //Добавляем ключ и ложим по нему пустую структуру(она нисколько не весит)
	}

	return set
}

func Task12() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	customSet := createSet(sequence)

	fmt.Println("Собственное множество:")
	for key := range customSet { //Пробегаемся по ключам т.к. в них хранятся значения
		fmt.Println(key)
	}
}
