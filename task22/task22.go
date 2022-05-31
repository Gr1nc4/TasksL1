package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	bigIntA := big.NewInt(104857600)
	bigIntB := big.NewInt(104857600)
	c := big.NewInt(0)

	fmt.Println("Результат умножения: ", c.Mul(bigIntA, bigIntB))
	fmt.Println("Результат деления: ", c.Div(bigIntA, bigIntB))
	fmt.Println("Результат сложения: ", c.Add(bigIntA, bigIntB))
	fmt.Println("Результат вычитания: ", c.Sub(bigIntA, bigIntB))
}
