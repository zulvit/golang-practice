package main

import (
	"fmt"
)

// Обмен с использованием арифметического метода
func swapArithmetic(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

// Обмен с использованием побитового XOR
func swapXOR(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

// Обмен с использованием множественного присваивания
func swapMultipleAssignment(a, b int) (int, int) {
	return b, a
}

// Обмен с использованием указателей
func swapPointers(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 10
	b := 20

	// Арифметический метод
	a, b = swapArithmetic(a, b)
	fmt.Println("После арифметического обмена: a =", a, "b =", b)

	// Метод XOR
	a, b = swapXOR(a, b)
	fmt.Println("После XOR обмена: a =", a, "b =", b)

	// Множественное присваивание
	a, b = swapMultipleAssignment(a, b)
	fmt.Println("После множественного присваивания обмена: a =", a, "b =", b)

	// Метод с использованием указателей
	swapPointers(&a, &b)
	fmt.Println("После обмена с использованием указателей: a =", a, "b =", b)
}
