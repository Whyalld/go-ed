package main

import (
	"fmt"
	"time"
)

// func main() {
// 	nums := make(chan int)
// 	squares := make(chan int)

// 	go func() {
// 		defer close(nums)

// 		for i := 1; i <= 10; i++ {
// 			nums <- i
// 			time.Sleep(300 * time.Millisecond)
// 		}
// 	}()

// 	go func() {
// 		defer close(squares)

// 		for n := range(nums) {
// 			squares <- n * n
// 		}
// 	}()

// 	for n := range squares {
// 		fmt.Println(n)
// 	}
// }

// type Result struct {
// 	Number int
// 	Square int
// }

// func generateNumbers(numbers chan<- int) {
// 	defer close(numbers)

// 	for i := 1; i <= 10; i++ {
// 		numbers <- i
// 		time.Sleep(300 * time.Millisecond)
// 	}
// }

// func calculateSquares(numbers <-chan int, results chan<- Result) {
// 	defer close(results)

// 	for num := range numbers {
// 		res := Result{Number: num, Square: num * num}
// 		results <- res
// 	}
// }

// func main() {
// 	numbers := make(chan int)
// 	results := make(chan Result)

// 	go generateNumbers(numbers)
// 	go calculateSquares(numbers, results)

// 	for n := range results {
// 		fmt.Printf("Квадрат числа %d = %d\n", n.Number, n.Square)
// 	}
// }

const workerCount = 3

type Result struct {
	WokerID int
	Number  int
	Square  int
}

func generateNumbers(userInput int, numbers chan<- int) {
	defer close(numbers)

	for i := 1; i <= userInput; i++ {
		numbers <- i
	}
}

func calculateSquares(
	wokerID int,
	numbers <-chan int,
	results chan<- Result,
	done chan<- struct{},
) {
	// Перед выходом сообщаем, что работник завершился
	defer func() {
		done <- struct{}{}
	}()

	for number := range numbers {
		// Имитируем продолжительные вычисления
		time.Sleep(300 * time.Millisecond)

		results <- Result{
			WokerID: wokerID,
			Number:  number,
			Square:  number * number,
		}
	}
}

func main() {
	var userInput int

	fmt.Print("Введите число: ")

	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("Ошибка: необходимо ввести целое число")
		return
	}

	if userInput <= 0 {
		fmt.Println("Ошибка: число должно быть больше нуля")
		return
	}

	numbers := make(chan int)
	results := make(chan Result)
	done := make(chan struct{})

	go generateNumbers(userInput, numbers)

	for workerID := 1; workerID <= workerCount; workerID++ {
		go calculateSquares(workerID, numbers, results, done)
	}

	// Эта горутина ждет завершения всех работников
	go func() {
		for i := 0; i < workerCount; i++ {
			<-done
		}

		close(results)
	}()

	// main получает и печатает готовые результаты
	for result := range results {
		fmt.Printf(
			"Работник %d: квадрат числа %d = %d\n",
			result.WokerID, result.Number, result.Square)
	}
}
