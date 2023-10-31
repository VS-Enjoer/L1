package Solution

import (
	"fmt"
	"sync"
)

/*
Задание 18:
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	mu    sync.Mutex //Добавляем мьютекс для корректной записи значения в переменную value
	value int
	wg    sync.WaitGroup
}

func (c *Counter) Increment() {
	c.mu.Lock() //Лочим горутины для безопасной записи в перменную value
	c.value++
	c.mu.Unlock() //Анлочим для того чтобы другие горутины тоже могли инкрементировать переменную
}

func Task18() {
	count := Counter{}

	GoRutines := 70

	//Увеличиваем счетчик горутин
	count.wg.Add(GoRutines)

	//Запускаем горутины и вызываем у них инкремент и после завершаем горутину
	for i := 0; i < GoRutines; i++ {
		go func() {
			count.Increment()
			count.wg.Done()
		}()
	}
	//Ждем завершение работы горутин
	count.wg.Wait()
	fmt.Print(count.value)
}
