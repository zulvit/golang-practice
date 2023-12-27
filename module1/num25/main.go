package main

import (
	"fmt"
	"time"
)

// Sleep функция приостанавливает выполнение программы на указанное количество секунд.
func Sleep(seconds int) {
	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)
}

func main() {
	fmt.Println("Начало работы программы")
	Sleep(3) // Приостанавливаем выполнение программы на 3 секунды
	fmt.Println("Прошло 3 секунды")
}
