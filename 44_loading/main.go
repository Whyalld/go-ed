package main

import (
	"fmt"
	"math"
	"time"
)

func showLoading(duration time.Duration) error {
	if !validDimension(float64(duration)) {
		return fmt.Errorf("Некорректное число: %v", duration)
	}

	symbols := `-\|/`
	delay := 25 * time.Millisecond
	stop := time.Now().Add(duration)

	for time.Now().Before(stop) {
		for _, symb := range symbols {
			if !time.Now().Before(stop) {
				break
			}

			fmt.Printf("\r%c", symb)
			time.Sleep(delay)
		}
	}

	fmt.Print("\rГотово!\n")
	return nil

	// for i := 1; i <= int(num)*10; {
	// 	for _, r := range symbols {
	// 		fmt.Printf("\r%c", r)
	// 		time.Sleep(delay)
	// 	}
	// 	i++
	// }
	// fmt.Print("\rFinish!\n")
	// return nil

}

func main() {
	var duration time.Duration
	fmt.Print("Введите число секунд: ")
	fmt.Scan(&duration)

	start := time.Now()

	if err := showLoading(duration * time.Second); err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("time:", time.Since(start).Seconds())
}

func validDimension(value float64) bool {
	return value >= 0 &&
		!math.IsNaN(value) &&
		!math.IsInf(value, 0)
}
