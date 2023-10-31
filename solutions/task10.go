package Solution

import (
	"fmt"
	"sort"
)

/*
Задание 10:
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

*/

func groupTemperatures(temperatures []float64) map[int][]float64 {
	sort.Float64s(temperatures) //Сортируем массив

	groupedTemperatures := make(map[int][]float64) //Создание мапы для записи значений по ключу

	for _, temp := range temperatures {
		key := int(temp/10) * 10                    // Разделение значений температур на диапазоны по 10 градусов
		if _, ok := groupedTemperatures[key]; !ok { // Проверка наличия ключа с соответствующим диапазоном в карте
			groupedTemperatures[key] = []float64{temp} // Если ключа нет, создаем новую запись в карте с найденным диапазоном
		} else {
			groupedTemperatures[key] = append(groupedTemperatures[key], temp) // Если ключ уже существует, добавляем значение температуры к существующему диапазону
		}
	}

	return groupedTemperatures
}

func Task10() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	grouped := groupTemperatures(temperatures)

	for key, value := range grouped {
		fmt.Printf("%d: %v\n", key, value)
	}
}
