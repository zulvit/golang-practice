package main

import (
	"fmt"
	"reflect"
)

func determineType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Тип int: %v\n", v)
	case string:
		fmt.Printf("Тип string: %v\n", v)
	case bool:
		fmt.Printf("Тип bool: %v\n", v)
	case chan int:
		fmt.Printf("Тип channel: %v\n", reflect.TypeOf(v))
	default:
		fmt.Printf("Неизвестный тип: %v\n", reflect.TypeOf(v))
	}
}

func main() {
	var myInt int = 5
	var myString string = "Привет"
	var myBool bool = true
	var myChan chan int = make(chan int)

	determineType(myInt)
	determineType(myString)
	determineType(myBool)
	determineType(myChan)
}
