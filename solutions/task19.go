package Solution

import "fmt"

/*
Задание 19:
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode
*/

func NextLvlStrginger(input string) string {
	runeArray := []rune(input) //конвертируем в слайс рун полученную строку(т.к. могут прилететь символы юникод)

	result := make([]rune, len(runeArray)) //Создаем слайс рун для заполнения, чтобы вернуть уже итоговый слайс

	for i := 0; i < len(result); i++ { //Переставляем элементы в обратном порядке
		result[i] = runeArray[len(runeArray)-1-i]
	}

	return string(result) //Конвертируем в строку и возвращаем строку
}

func Task19() {
	inputString := "☀ ❤文🙂"

	fmt.Println(NextLvlStrginger(inputString))
}
