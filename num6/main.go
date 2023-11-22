package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	stopChan := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	quitFlag := false
	var wg sync.WaitGroup
	timeout := time.After(time.Duration(10) * time.Second)
	wg.Add(1)
	go worker(stopChan, ctx, &quitFlag, &wg, timeout)

	// Различные способы остановки горутины
	time.Sleep(3 * time.Second)
	stopChan <- struct{}{}

	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(3 * time.Second)
	quitFlag = true

	wg.Wait()
}

// Функция, которую я останавливаю всеми возможными способами
func worker(stopChan <-chan struct{}, ctx context.Context, quitFlag *bool,
	wg *sync.WaitGroup, timeout <-chan time.Time) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Println("Остановка через канал")
			return
		case <-ctx.Done():
			fmt.Println("Остановка через контекст")
			return
		case <-timeout:
			fmt.Println("Остановка по таймауту")
			return
		default:
			if *quitFlag {
				fmt.Println("Остановка через флаг")
				return
			}
			fmt.Println("Работаю...")
			time.Sleep(time.Second)
		}
	}
}
