package main

import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "A"
	}
	return v
}

/*
Мы создаём весьма большую строку, однако, после используем первые 100 элементов, поэтому
я решил использовать в новой реализации copy, чтобы создать новую строку для экономии места
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}
*/

func someFunc() {
	v := createHugeString(1 << 10)
	// Использование copy для создания новой строки
	tempString := make([]byte, 100)
	copy(tempString, v[:100])
	justString = string(tempString)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
