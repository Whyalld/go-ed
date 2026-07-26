package main

import "fmt"

func main() {
	var number int

	fmt.Print("Введите число: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Ошибка: нужно ввессти целое число")
		return
	}

	fmt.Println("Вы ввели:", number)
}