package main

import (
	"fmt"
	"time"
)

func main() {
	resultChan := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		result := 10 + 20
		resultChan <- result
	}()

	fmt.Println("Выполняем другую работу")

	result := <-resultChan
	fmt.Println(result) // 30
}
