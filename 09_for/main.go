package main

import "fmt"

// В go нету while, loop, есть только for

func main9() {
	// Можно не указывать второй параметр для функции, еслии исп только первый
	var list = [...]string{
		"молоко",
		"чипсы",
		"кола",
	}
	for index := range list {
		fmt.Printf("%d. %s\n", index+1, list[index])
	}
}

func mainY() {
	// For используется для обычных итераций
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func mainA() {
	// Можно опускать некоторые блоки
	for i := 0; i < 10; {
		i++
		fmt.Println(i)
	}
}

func main3() {
	// Бесконечный цикл
	for i := 1; ; i++ {
		println(i)
	}
}

func main4() {
	// Бесконеный цикл (реально)
	var i int = 1
	for {
		if i > 100 {
			break
		}

		fmt.Println(i)
		i++
	}
}

func main11() {
	var arr [3]int
	// fillArray(arr)
	arr = fillArray1(arr)
	fmt.Println(arr)
}

func fillArray(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println("fillArray():", arr)
}

func fillArray1(arr [3]int) [3]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return arr
}
