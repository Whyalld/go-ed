package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // время старта
	done := make(chan struct{})

	fmt.Printf("Старт: %s\n", start.Format(time.RFC3339))

	go CalculateSomething(1000, done)

	go CalculateSomething(2000, done)

	<-done
	<-done

	fmt.Printf("Время выполнения программы: %s\n", time.Since(start))
}


func CalculateSomething(n int, done chan <- struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	
	t := time.Now()

	result := 0
	for i := 0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Microsecond * 3)
	}

	fmt.Printf("Рузультат: %d; Прошло времени: %s\n", result, time.Since(t))
}

