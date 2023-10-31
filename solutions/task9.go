package Solution

import "fmt"

/*
Задание 9:
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func Task9() {
	numbers := []int{1, 2, 3, 4, 5}

	input := make(chan int)  //Канал для записи
	output := make(chan int) //Канал для вывода

	// Горутина для записи чисел в первый канал
	go func() {
		defer close(input)
		for _, num := range numbers {
			input <- num
		}
	}()

	// Горутина для удвоения чисел из первого канала и записи результатов во второй канал
	go func() {
		defer close(output)
		for num := range input {
			output <- num * 2
		}
	}()

	// Вывод результатов из второго канала в stdout
	for doubled := range output {
		fmt.Println(doubled)
	}
}
