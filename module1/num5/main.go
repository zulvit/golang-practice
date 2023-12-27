package main

import (
	"fmt"
	"time"
)

func main() {
	// Канал для передачи данных
	dataChan := make(chan int)
	go sendData(dataChan)
	go readData(dataChan)

	// Завершение программы через N секунд
	N := 5
	timer := time.NewTimer(time.Duration(N) * time.Second)

	<-timer.C
	fmt.Println("Программа звершена")
}

// Отправка данных
func sendData(ch chan int) {
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

// Чтение данных
func readData(ch chan int) {
	for i := 1; i < 10; i++ {
		value := <-ch
		fmt.Printf("Прочитано из канала %d\n", value)
	}
}
