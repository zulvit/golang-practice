package main

import (
	"fmt"
	"sync"
)

func main() {
	safeMap := SafeMap{data: make(map[string]int)}
	var wg sync.WaitGroup
	// Запись значений в map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("Ключ%d", n)
			safeMap.Set(key, n)
		}(i)
	}

	wg.Wait()

	fmt.Println(safeMap)
}

type SafeMap struct {
	data map[string]int
	mu   sync.Mutex
}

// Set установит значние
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}
