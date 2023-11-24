package main

import (
	"fmt"
)

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создание множества
	set := make(map[string]bool)

	// Добавление каждой строки из последовательности в множество
	for _, str := range strings {
		set[str] = true
	}

	// Вывод множества
	for key := range set {
		fmt.Println(key)
	}
}
