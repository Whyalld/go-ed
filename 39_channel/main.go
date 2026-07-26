package main

import "fmt"
import "time"

func main() {
	nums := make(chan int)
	squares := make(chan int)
	
	go func() {
		defer close(nums)

		for i := 1; i <= 10; i++ {
			nums <- i
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		defer close(squares)

		for n := range(nums) {
			squares <- n * n
		}
	}()

	for n := range squares {
		fmt.Println(n)
	}
}