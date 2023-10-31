package Solution

import "fmt"

/*
	Задание 3:
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов с использованием конкурентных вычислений.
*/

func NextSolution(chanel chan int, MyArray []int) {
	defer close(chanel)
	//Перебираем элементы массива и ложим в канал их значение в квадрате
	for _, elem := range MyArray {
		chanel <- elem * elem

	}

}

func Task3() {
	MyArray := []int{2, 4, 6, 8, 10}

	//Создание буферизованного канала чтобы потом перебрать значения из канала и узнать сумму
	chanel := make(chan int, len(MyArray))
	var result int

	go NextSolution(chanel, MyArray)

	//Перебираем значения из канала
	for elem := range chanel {
		result += elem
	}

	fmt.Println(result)
}
