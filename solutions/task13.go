package Solution

import "fmt"

/*
Задание 13:
Поменять местами два числа без создания временной переменной
*/

func SwapNumbers(a, b *int) {
	*a, *b = *b, *a
}

func Task13() {
	a, b := 23, 12

	fmt.Println(a, b)

	SwapNumbers(&a, &b)

	fmt.Println(a, b)
}
