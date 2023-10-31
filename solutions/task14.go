package Solution

import (
	"fmt"
	"reflect"
)

/*
Задание 14:
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

func getType(i interface{}) string {
	// Получаем тип переменной во время выполнения
	t := reflect.TypeOf(i)

	// Возвращаем имя типа
	return t.String()
}

func Task14() {
	var num int = 5
	var text string = "Hello"
	var check bool = true
	var ch chan int

	// Используем функцию getTypeName для определения типа переменной во время выполнения
	fmt.Println("Тип переменной 'num':", getType(num))
	fmt.Println("Тип переменной 'text':", getType(text))
	fmt.Println("Тип переменной 'check':", getType(check))
	fmt.Println("Тип переменной 'ch':", getType(ch))
}
