package main

import "fmt"

func main13_2() {
	age := 30 // Обычная переменная

	// &age берет адрес ячейки памяти (например, 0xc0000140a8)
	pointerToAge := &age

	fmt.Println(pointerToAge) // Выведет адрес в памяти: 0xc0000140a8
	fmt.Println(*pointerToAge) // Выведет значение по этому адресу: 25

	// Меняем значение через указатель
	*pointerToAge = 18

	fmt.Println(age) // Выведет 30! Оригинальная переменная изменилась
}