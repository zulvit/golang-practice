package main

import (
	"fmt"
)

func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...) //использовал распаковку среда
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	i := 2 // Индекс элемента, который нужно удалить

	// Удаление элемента
	mySlice = remove(mySlice, i)

	fmt.Println(mySlice)
}
