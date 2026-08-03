package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите возраст: ")

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
		os.Exit(1)
	}

	age, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Некорректный возраст")
		os.Exit(1)
	}

	fmt.Println("Ваш возраст: ", age)
}