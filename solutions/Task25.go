package Solution

import (
	"fmt"
	"time"
)

/*
Задание 25:
Реализовать собственную функцию sleep
*/

func sleep(duration time.Duration) {
	endTime := time.Now().Add(duration) // Вычисляем время окончания ожидания добавляя к текущему времени заданное время
	for time.Now().Before(endTime) {
	} //Проверяем вышло ли время с помощью метода before передавая туда значение до какого времени мы спим
}

func Task25() {
	fmt.Println("Начало")
	sleep(2 * time.Second) // Ожидаем 2 секунды
	fmt.Println("Задержка закончена")
}
