package main

import "fmt"

func main7() {
	// Make создает slice из n-го количества элементов
	slice := make([]int, 5)      // len = 5, cap = 5
	slice2 := make([]int, 0, 10) // len = 0, cap = 10

	// Copy копирует срезы/массивы, но не связывает с ними
	list := [...]int{1, 2, 4, 5}
	data := make([]int, 4) // Сначала выделяем память для копирования через make
	copy(data, list[:])    // [:] нужно т.к list - массив

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(list)
	fmt.Println(data)

}
