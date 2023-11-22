package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	square := calcSumOfSquare(numbers)
	fmt.Println(square)
}

func calcSumOfSquare(num []int) int {
	var sum int
	var waitGroup sync.WaitGroup // Использовал WaitGroup того чтобы подождать выполнения
	var mutex sync.Mutex         // Использовал мьютекс для блокировки

	for _, v := range num {
		waitGroup.Add(1)
		go func(v int) {
			defer waitGroup.Done()
			square := v * v
			mutex.Lock()
			sum += square
			mutex.Unlock()
		}(v)
	}
	waitGroup.Wait()
	return sum
}
