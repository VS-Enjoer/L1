package Solution

import (
	"fmt"
	"time"
)

/*
	Задание 5:
	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой
	стороны канала — читать. По истечению N секунд программа должна завершаться.

*/

func WriteSolution(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i                 // Отправляем значения в канал
		time.Sleep(time.Second) // Ждем 1 секунду перед отправкой следующего значения
	}
	close(ch) // Закрываем канал после отправки всех значений
}

func Task5() {
	ch := make(chan int)

	go WriteSolution(ch) // Запускаем функцию отправки значений в канал в отдельной горутине

	timer := time.NewTimer(10 * time.Second) // Создаем таймер на 10 секунд

	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("Получено из канала:", val)
		case <-timer.C:
			fmt.Println("Время истекло")
			return
		}
	}
}
