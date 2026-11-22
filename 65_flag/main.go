package main

import (
	"flag"
	"fmt"
)

func main() {
	// Определяем флаг --age или -age со значением по умолчанию 0
	age := flag.Int("age", 0, "Ваш возраст для авторизации")
	
	// Парсим переданные аргументы
	flag.Parse()

	if *age <= 0 {
		fmt.Println("Ошибка: укажите корректный возраст через флаг -age")
		return
	}

	fmt.Println("Ваш возраст:", *age)
}
