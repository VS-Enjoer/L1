package Solution

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Задание 4:
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func CntrlCSolution(ctx context.Context, chanel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		//Проверка завершения контекста
		case <-ctx.Done():
			return

			//Вывод данных из канала
		case element, ok := <-chanel:
			if !ok {
				return
			}
			fmt.Println("Полученно:", element)
		}
	}
}

func Task4() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Создание не буферизованного канала потому что он хранит 1 значение и выводит его в терминал, после берет новое значение
	chanel := make(chan int)
	GoRutCount := 5 // кол-во горутин
	var wg sync.WaitGroup

	for i := 0; i < GoRutCount; i++ {
		wg.Add(1)
		go CntrlCSolution(ctx, chanel, &wg)
	}

	// Слушаем сигналы SIGINT (Ctrl+C)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println("Получен сигнал SIGINT, инициируем завершение...")
		cancel() // Отправляем сигнал завершения при нажатии Ctrl+C
	}()

	// Отправляем данные в канал пока не будет Ctrl+C
	go func() {
		for i := 0; ; i++ {
			select {
			case chanel <- i: // Пишем данные в канал
			case <-ctx.Done():
				fmt.Println("Главный поток получил сигнал остановки записи в канал.")
				close(chanel) // Закрываем канал по сигналу главного потока
				return
			}
		}
	}()

	wg.Wait() // Ждем завершения всех горутин
	fmt.Println("Горутины завершили работу.")
}
