package Solution

import (
	"fmt"
	"math/big"
)

/*
Задание 22:
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных
a,b, значение которых > 2^20.
*/

func Task22() {
	//Использую big.Int потому как значения будут судя по всему очень большие. big.Int расширяется динамически (пока хватает оперативной памяти)
	a := big.NewInt(712512341)
	b := big.NewInt(512412512)

	proizv := new(big.Int).Mul(a, b)
	chastn := new(big.Int).Div(a, b)
	sum := new(big.Int).Add(a, b)
	minus := new(big.Int).Sub(a, b)

	fmt.Printf("Сумма a и b = %s\n", sum)
	fmt.Printf("Разность a и b = %s\n", minus)
	fmt.Printf("Произведение a и b = %s\n", proizv)
	fmt.Printf("Частное a и b = %s\n", chastn)
}
