package main

import (
	"fmt"
	"time"
)

// or функция объединяет несколько каналов в один.
var or func(channels ...<-chan interface{}) <-chan interface{}

func init() {
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		default:
			orDone := make(chan interface{})
			go func() {
				defer close(orDone)
				switch len(channels) {
				case 2:
					select {
					case <-channels[0]:
					case <-channels[1]:
					}
				default:
					select {
					case <-channels[0]:
					case <-channels[1]:
					case <-channels[2]:
					case <-or(append(channels[3:], orDone)...):
					}
				}
			}()
			return orDone
		}
	}
}

// sig функция создает канал, который закрывается после заданного времени.
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
