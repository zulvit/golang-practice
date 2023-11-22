package main

import (
	"fmt"
)

// setBit устанавливает i-й бит числа num в значение bitVal (1 или 0)
func setBit(num int64, i uint, bitVal int) int64 {
	if bitVal == 1 {
		// Установка бита в 1
		num |= 1 << i
	} else {
		// Установка бита в 0
		num &= ^(1 << i)
	}
	return num
}

func main() {
	var num int64 = 0
	var pos uint = 2
	var bitVal int = 1

	// Установка бита
	num = setBit(num, pos, bitVal)
	fmt.Printf("Результат: %b\n", num)

	// Изменение бита
	bitVal = 0
	num = setBit(num, pos, bitVal)
	fmt.Printf("Результат: %b\n", num)
}
