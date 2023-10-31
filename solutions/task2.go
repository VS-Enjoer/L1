package Solution

import (
	"fmt"
	"sync"
)

/*
	Задание 2:
	Написать программу, которая конкурентно рассчитает значение квадратов чисел
	взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func FirstSolution(wg *sync.WaitGroup, MyArray []int) {

	for _, elem := range MyArray { //Перебираем элементы массива, получаем квадрат и выводим в терминал
		wg.Add(1)

		go func(e int) {
			defer wg.Done()
			result := e * e
			fmt.Println(result)

		}(elem)
	}

}

func Task2() {
	MyArray := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	FirstSolution(&wg, MyArray)

	wg.Wait() //ожидание завершения потоков
}
