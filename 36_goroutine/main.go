package main

import "fmt"

func main() {
	input := make(chan int)
	output := make(chan int)

	go func() {
		defer close(output)
		for val := range input {
			output <- val
		}
	}()

	go func() {
		input <- 10
		input <- 20
		input <- 30
		close(input)
	}()

	fmt.Println(<-output)
	fmt.Println(<-output)
	fmt.Println(<-output)
}
