package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	// Задаем значения a и b больше 2^20
	a.SetString("1048576", 10) // 2^20
	b.SetString("2097152", 10) // 2^21

	// Умножение
	mul := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s\n", mul.String())

	// Деление
	div := new(big.Int).Div(b, a)
	fmt.Printf("Деление: %s\n", div.String())

	// Сложение
	add := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s\n", add.String())

	// Вычитание
	sub := new(big.Int).Sub(b, a)
	fmt.Printf("Вычитание: %s\n", sub.String())
}
