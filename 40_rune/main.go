package main

import (
	"fmt"
	"time"
)

func main() {
	var char rune = '💁'
	fmt.Printf("%c\n", char)

	for i := 1; i <= 10; i++ {
		fmt.Printf("\r%d", i)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println()
}