package main

import (
	"fmt"
	"time"
	// "sync"
)

// type factorialResult struct {
// 	value int
// 	err error
// }

func factorial(n int) int {
	time.Sleep(1000 * time.Millisecond)
	
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func spinner(done <-chan struct{}, stopped chan<- struct{}) {
	defer close(stopped)

	symbols := []rune(`-\|/`)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	position := 0

	for {
		select {
		case <-done:
			fmt.Print("\r				\r")
			return

		case <-ticker.C:
			fmt.Printf("\rВычисление... %c", symbols[position])
			position = (position + 1) % len(symbols)
		}
	}
}

func main() {
	result := make(chan int)
	done := make(chan struct{})
	stopped := make(chan struct{})

	go spinner(done, stopped)

	go func() {
		result <- factorial(10)
	}()

	value := <-result

	close(done)
	<-stopped

	fmt.Printf("Результат: %d\n", value)
}