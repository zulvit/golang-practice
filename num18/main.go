package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait() // ожидаю пока все горутины отработают перед завершением
	fmt.Println("Итоговое значение счётчика:", counter.value)
}

// Counter - структура для итерации в многопоточной среде
type Counter struct {
	value int
	mutex sync.Mutex
	//Для создания структуры, которая сможет безопасно итерироваться
	//я использую механизм блокировки из пакета sync - Mutex
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

func (c *Counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock() // автоматическая разблокировка при выходе из метода
	return c.value
}
