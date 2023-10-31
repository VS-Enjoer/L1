package Solution

import (
	"fmt"
)

/*
	Задание 1:
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)
*/

// Основная структура

type Human struct {
	Name   string
	Age    int
	Gender string
}

// Структура которая наследуется от структуры Human и так же добавлено дополнительное поле ActionNow(Действие сейчас)

type Action struct {
	Human
	ActionNow string
}

// Реализация метода у структуры Human (так же будет доступен при вызове у наследников)

func (h *Human) Getinfo() {
	fmt.Printf("Привет, мое имя %s, мне %d лет, мой пол: %s.\n", h.Name, h.Age, h.Gender)
}

// Реализация метода у структуры Action

func (a *Action) Run() {
	fmt.Printf("Привет, мое имя %s, мне %d лет, мой пол: %s и я сейчас мегабыстро %s.\n", a.Name, a.Age, a.Gender, a.ActionNow)
}

func Task1() {

	Dima := Action{

		Human:     Human{Name: "Дима", Age: 27, Gender: "мужской"},
		ActionNow: "бегу",
	}
	//Методы вызываются у структуры Action и т.к. выше она была унаследованна, вызывается метод Human и метод Action для демонстрации
	Dima.Getinfo()
	Dima.Run()

}
