package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	waitGroup := sync.WaitGroup{} //Использовал WaitGroup того чтобы подождать выполнения

	for _, v := range numbers {
		waitGroup.Add(1)
		go func(v int) {
			defer waitGroup.Done()
			calcSquare(v)
		}(v)
	}

	waitGroup.Wait()
}

func calcSquare(num int) {
	fmt.Println(num * num)
}
