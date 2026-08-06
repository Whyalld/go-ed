package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Введите число: ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка: ", err)
		return
	}

	num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "пожалуйста, введите число")
	}

	switch {
	case num < 10: fmt.Printf("Number %v is less than 10\n", num)
	case num > 10: fmt.Printf("Number %v is greater than 10\n", num)
	default: fmt.Println("Number is equal to 10")
	}

}
