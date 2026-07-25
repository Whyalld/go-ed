package main

import (
	"fmt"
	"time"
	"math"
)

func download(num float64) error{
	alph := `-\|/`
	if !validDimension(num) {
		return fmt.Errorf("Некорректное число: %d", int(num))
	}
	for i := 1; i <= int(num) * 10; {
		for _, r := range(alph) {
			fmt.Printf("\r%c", r)
			time.Sleep(time.Millisecond * 25)
		}
		i++
	}
	fmt.Print("\r Finish!\n")
	return nil
}

func main() {
	start := time.Now()
	download(5)
	fmt.Println(time.Since(start).Seconds())
}

func validDimension(value float64) bool {
	return value >= 0 &&
		!math.IsNaN(value) &&
		!math.IsInf(value, 0)
}