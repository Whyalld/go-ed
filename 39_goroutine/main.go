package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() // время старта

	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	go calculateSomething(1000)

	go calculateSomething(2000)

	time.Sleep(8 * time.Second)

	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))
}

func calculateSomething(n int) {
	t := time.Now()

	result := 0
	for i := 0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Microsecond * 3)
	}

	fmt.Printf("Рузультат: %d; Прошло времени: %s\n", result, time.Since(t))
}
