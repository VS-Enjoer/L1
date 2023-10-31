package Solution

import "fmt"

/*
Задание 8:
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(num int64, i int, value int) int64 {
	// Устанавливаем i-й бит в 1
	if value == 1 {
		num = num | (1 << i)
	} else if value == 0 {
		// Устанавливаем i-й бит в 0
		num = num &^ (1 << i)
	}
	return num
}

func Task8() {
	var num int64 = 27 // Пример числа
	i := 3             // Индекс бита
	bitValue := 0      // Устанавливаем значение бита (0 или 1)

	// Устанавливаем i-й бит в указанное значение
	result := setBit(num, i, bitValue)

	fmt.Printf("Результат: %d\n", result)
}
