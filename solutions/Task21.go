package Solution

import (
	"fmt"
)

/*
Задание 21:
Реализовать паттерн «адаптер» на любом примере.


Как я понял паттерн адаптер нам говорит о том что у нас в программе могут быть 2 несовместимые структуры и нужно сделать
дополнительный адаптер который будет преобразовывать одну структуру в другую. Чтобы программа могла работать с ними.

В моем решении мы на вход получаем json файл и преобразовываем его в xml для дальнейшей работы
*/

// XMLData - структура для работы с XML

type XMLData struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

// JSONData - структура для работы с JSON

type JSONData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// XMLAdapter - адаптер для преобразования JSON в XML

type XMLAdapter struct {
	JSONData
}

// XmlDataFromJSON - метод адаптера для преобразования JSON в XML

func (adapter *XMLAdapter) XmlDataFromJSON() XMLData {
	return XMLData{
		Name: adapter.Name,
		Age:  adapter.Age,
	}
}

// Вывод данных из формата xml

func ProcessXMLData(data XMLData) {
	fmt.Printf("Результат работы Xml: Имя - %s, Возраст - %d лет\n", data.Name, data.Age)
}

// Вывод данных из формата json

func ProcessJSONData(data JSONData) {
	fmt.Printf("Результат работы Json: Имя - %s, Возраст - %d лет\n", data.Name, data.Age)
}

func Task21() {
	// Данные в формате JSON
	jsonData := JSONData{Name: "Олег", Age: 30}

	// Создание адаптера
	adapter := XMLAdapter{jsonData}

	// Преобразование JSON в XML через адаптер
	xmlData := adapter.XmlDataFromJSON()

	//Работа Json метода
	ProcessJSONData(jsonData)

	// Работа Xml метода
	ProcessXMLData(xmlData)
}
