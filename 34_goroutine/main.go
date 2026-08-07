package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 10
	}()

	ch2 := make(chan string)
	go func() {
		ch2 <- "hello"
	}()

	num := <-ch1
	fmt.Println(num)

	select {
	case v := <-ch1:
		fmt.Println("channel 1 sends", v)
	case v := <-ch2:
		fmt.Println("channel 2 sends", v)
	default:
		fmt.Println("neither channel was ready")
	}
}
