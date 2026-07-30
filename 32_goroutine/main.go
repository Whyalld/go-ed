package main

import (
	"fmt"
	"time"
)

func main() {
	timerChan := make(chan time.Time)

	go func() {
		time.Sleep(2 * time.Second)
		timerChan <- time.Now()
	}()

	completeAt := <-timerChan
	fmt.Println(completeAt)
}