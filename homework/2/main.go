// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	err := do(0)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func do(num int) error {
// 	if num < 1 {
// 		return errors.New("Некорректное число")
// 	}
// 	for i := range num {
// 		fmt.Println(1 << (i + 1))
// 	}
// 	return nil
// }

package main

import (
	"errors"
	"fmt"
	"os"
)

// ErrInvalidNumber выносится на уровень пакета, если ошибку нужно проверять в других местах.
var ErrInvalidNumber = errors.New("число должно быть натуральным (больше 0)")

func main() {
	var userInput int
	fmt.Print("Введите натуральное число: ")

	if _, err := fmt.Scanln(&userInput); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода: ", err)
		os.Exit(1)
	}

	fmt.Println("Result:")

	// В Go принято обрабатывать ошибку сразу в блоке if (if с инициализацией)
	if err := printPowersOfTwo(userInput); err != nil {
		// Ошибки в прод-коде часто выводят в поток stderr, а не stdout
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
}

// 1. Имя функции изменено на говорящее.
// 2. Вводной аргумент задокументирован.
func printPowersOfTwo(num int) error {
	if num < 1 {
		return ErrInvalidNumber
	}

	for i := 0; i < num; i++ {
		fmt.Println(1 << (i + 1))
	}
	return nil
}
