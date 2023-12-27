package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{1, 634, 215, 272, 39, 0, 4, 123, 333}
	sort.Ints(slice) //Сортирую слайс для бинарного поиска
	fmt.Println(slice)
	target := 4

	// Использование sort.Search для бинарного поиска
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
		//Возвращает true, если искомый элемент находится на текущем индексе или правее его.
		//Возвращает false, если искомый элемент находится левее текущего индекса.
	})

	// Проверка, найдено ли число
	if index < len(slice) && slice[index] == target {
		fmt.Printf("Найдено число %d на позиции %d.\n", target, index)
	} else {
		fmt.Println("Число не найдено.")
	}
}
