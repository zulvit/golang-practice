package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	inputChan := make(chan int)
	outputChan := make(chan int)

	// Горутина для умножения числа на 2
	go func() {
		for i := range inputChan {
			outputChan <- i * 2
		}
		close(outputChan)
	}()

	// Горутина для отправки чисел в inputChan
	go func() {
		for _, v := range numbers {
			inputChan <- v
		}
		close(inputChan)
	}()

	for result := range outputChan {
		fmt.Println(result)
	}
}
