package main

import (
	"fmt"
	// "time"
)

func main() {
	numbers := make(chan int)

	go generateNumbers(1000, numbers)

	// for number := range numbers {
	// 	fmt.Println(number)
	// }

	for {
		number, ok := <-numbers

		fmt.Println(number, ok)

		if !ok {
			break
		}
	}
}

func generateNumbers(n int, res chan int) {
	defer close(res)

	for i := 0; i <= n; i++ {
		res <- i * 2
	}
}