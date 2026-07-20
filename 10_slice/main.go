package main

import "fmt"

func main5() {
	// Slice (срез) - тот же массив, но его длина динамическая
	var slice = []string{
		"Наушники",
		"Шнурки",
		"Брюки",
	}
	fmt.Println("Длина списка: ", len(slice))
	fmt.Println("Емокость списка: ", cap(slice))

	// for index, item := range slice {
	// 	fmt.Printf("%d. %s\n", index+1, item)
	// }

	slice = append(slice, "Купить нож")

	fmt.Println("Длина списка: ", len(slice))
	fmt.Println("Емокость списка: ", cap(slice))
}
