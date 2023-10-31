package Solution

import "fmt"

/*
Задания 20:
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».


Можно было воспользоваться модулем strings для выполнения задания а конкретно strings.Fields() для выделения слов в срез.
Но сделал со своим подходом, только кол-во пробелов на вывод будет 1 между словами, попробовал учесть еще что между слов
будет 2 и более пробелов на входе но на вывод этого не сделал а так можно бы было так же создать маппу и делать подсчет
пробелов. Может в глаза бросится наличие булеан перменной, но она необходима для работы с 2 и более пробелами в моей
функции.
*/

func ChangeString(input string) string {

	inputSlice := []rune(input)
	outputMap := make(map[int][]rune) //Мапа для хранения слов по ключу
	wordStarted := false              //Нужна для учета 2 и более пробелов(В плане чтобы SchetIndex не увеличивался если 2 пробела подряд)
	SchetIndex := 0

	for i := 0; i < len(input); i++ {
		if inputSlice[i] != ' ' { //проверяем что символ не является пробелом
			outputMap[SchetIndex] = append(outputMap[SchetIndex], inputSlice[i]) //Добавляем в мапу под нужным индексом символ
			wordStarted = true
		} else {
			if wordStarted { //Проверка чтобы не учитывать не нужные пробелы
				SchetIndex++        //Увеличиваем счетчик чтобы понимать сколько у нас слов
				wordStarted = false //Стоит ложь, чтобы на следующую итерацию не увеличивать счетчик при пробеле
			}
		}
	}
	resultSlice := make([]rune, 0) //Слайс на вывод

	for i := SchetIndex; i >= 0; i-- {

		resultSlice = append(resultSlice, outputMap[i]...) //Распаковка элемента мапы с ключем i в слайс
		if i > 0 {
			resultSlice = append(resultSlice, ' ') //Добавляем пробел для корректного вывода
		}
	}
	if resultSlice[0] == ' ' { //Убираем пробел если он был в конце input т.к. это увеличит счетчик и соответственно будет пустой элемент
		copy(resultSlice[0:], resultSlice[1:])
	}
	return string(resultSlice)
}

func Task20() {
	inputString := "    s  qwe adsf r qfqwerf     "
	fmt.Print(ChangeString(inputString))
}

// Кратко зарезюмирую, у нас по идее может прилететь пустая строка, проверки я не делал. Так же не сделал работу с юникод символами, не знаю нужна ли была она
