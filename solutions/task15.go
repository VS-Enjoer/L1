package Solution

import (
	"strings"
)

/*
Задание 15:
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.
*/

// Допустим у нас есть подобная функция для создания строки
func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	/* justString = v[:100] это не правильная реализация так как justString ссылается на строку v, которая очень большая,
	не рационально расходовать так память, потому как глобальная переменная всегда будет ссылаться на локальную большого
	размера, и она не будет очищена.

	Если кол-во элементов меняется при создании v. Тогда рационально было бы сделать проверку подобную
	if len(v) >= 100 {
		justString = string(v[:100])
	} else {
		justString = v
	}

	А если постоянное кол-во элементов у v, тогда достаточно написать так.
	*/
	justString = string(v[:100])

}

func Task15() {
	someFunc()
}
