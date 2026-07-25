package main

import (
	"fmt"
	"time"
)

func main() {
	symbols := `-\|/`
	
	for i := 0; i <= 10; i++ {
		for _, r := range(symbols) {
			fmt.Printf("\r%c", r)
			time.Sleep(time.Millisecond * 100)
		}

	}
	fmt.Print("\r")
	fmt.Println("Boo!")
	fmt.Scan()
}