package Solution

import (
	"fmt"
	"strings"
)

/*
Задание 26:
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func TrueOrFalse(str string) bool {
	str = strings.ToLower(str)
	result := []rune(str)

	//Ну тут пробегаемся по массиву rune преобразованного из строки и в случае повторения отправляем false иначе отправляем true
	for i := 0; i < len(result); i++ {
		for k := i + 1; k < len(result); k++ {

			if result[k] == result[i] {
				return false
			}
		}
	}

	return true
}

func Task26() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Print(TrueOrFalse(str1))
	fmt.Print(TrueOrFalse(str2))
	fmt.Print(TrueOrFalse(str3))
}
