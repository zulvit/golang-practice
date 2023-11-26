package main

import (
	"fmt"
	"strings"
)

func IsUnique(s string) bool {
	// Приводим строку к нижнему регистру для регистронезависимой проверки.
	s = strings.ToLower(s)

	// Создаем map для отслеживания уникальных символов.
	characters := make(map[rune]bool)

	// Перебираем символы в строке.
	for _, char := range s {
		// Если символ уже встречался, возвращаем false.
		if characters[char] {
			return false
		}
		// Иначе, помечаем символ как встреченный.
		characters[char] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("IsUnique(\"%s\"): %v\n", str1, IsUnique(str1))
	fmt.Printf("IsUnique(\"%s\"): %v\n", str2, IsUnique(str2))
	fmt.Printf("IsUnique(\"%s\"): %v\n", str3, IsUnique(str3))
}
