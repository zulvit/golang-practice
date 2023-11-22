package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Создание канала для передачи данных
	dataChannel := make(chan string)
	// Создание канала для завершения
	// как я понял, пустые структуры не занимаю памяти,
	// поэтому мы можем использовать их для сигналов
	done := make(chan struct{})

	var wg sync.WaitGroup

	// Количество воркеров
	numOfWorkers := 5
	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, done, &wg)
	}

	// Запись из главного потока каких-то данных
	go func() {
		defer close(dataChannel) //Закрытие канала по завершении работы
		for i := 1; i <= 10; i++ {
			dataChannel <- fmt.Sprintf("Данные %d", i)
		}
	}()

	// Обработка сигнала Ctrl+C для завершения работы.
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	select {
	case <-signalChannel:
		close(done)
		wg.Wait()
		fmt.Println("Программа завершена.")
	}
}

// Воркер, обладает id для удобства
func worker(id int, dataChannel <-chan string,
	done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Воркер %d начал работу.\n", id)
	for {
		select {
		case data, ok := <-dataChannel:
			if !ok {
				fmt.Printf("Воркер %d завершил работу\n", id)
				return
			}
			fmt.Printf("Воркер %d получил данные %s\n", id, data)
			return
		case <-done:
			fmt.Printf("Воркер %d завершил работу по сигналу\n", id)
			return
		}
	}
}
