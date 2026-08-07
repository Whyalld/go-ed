package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите число: ")

	var num int
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v, %v не является числом\n", err, num)
	}

	switch {
	case num < 10:
		fmt.Printf("Number %v is less than 10\n", num)
	case num > 10:
		fmt.Printf("Number %v is greater than 10\n", num)
	default:
		fmt.Println("Number is equal to 10")
	}
}
