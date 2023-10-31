package Solution

import (
	"fmt"
	"sync"
)

/*
Задание: 7
Реализовать конкурентную запись данных в map.
*/

func Task7() {
	data := make(map[int]int)
	var mutex sync.Mutex

	var wg sync.WaitGroup
	writeCount := 10

	for i := 0; i < writeCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// Блокируем доступ к мапе для записи
			mutex.Lock()
			data[index] = index * 2
			// Разблокируем доступ к мап
			mutex.Unlock()
		}(i) //Передаем следующий индекс в анонимную функцию
	}

	wg.Wait() //Ждем завершение горутин

	fmt.Println("Результаты записи в map:")
	for key, value := range data {
		fmt.Printf("%d: %d\n", key, value)
	}
}
