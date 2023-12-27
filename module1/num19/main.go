package main

import "fmt"

func main() {
	input := "Привет WB"
	output := reverseString(input)
	fmt.Printf("Исходная строка: %s\nПеревёрнутая строка: %s\n", input, output)
}

/*
Символы Unicode могут занимать больше одного байта.
В Go строки представлены последовательностью байтов,а не символов.
Поэтому для корректного переворачивания строки с символами Unicode лучше всего использовать руну,
*/
func reverseString(s string) string {
	runes := []rune(s)
	// Переворачивание слайса рун, одновременно иду с конца и начала
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
