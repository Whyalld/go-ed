package main

import "fmt"

func main() {
	// Массив - список фиксированной длины
	var toDoList = [3]string{
		"купить хлеб",
		"сделать уроки",
		"помыться",
	}

	fmt.Println("Количество элементов в списке:", len(toDoList))

	for index, item := range toDoList {
		fmt.Printf("%d) %s\n", index+1, item)
	}
}

func main2() {
	// Не обязательно заполнять сразу все эелементы массива
	var toDoList1 = [4]string{"Выйти погулять"}
	toDoList1[3] = "Покуралесить"
	fmt.Println(toDoList1)
}

func main13() {
	// Можно вот так сделать
	var array = [...]int{
		1, 23, 32, 123, 123, 123, 213, 32132143, 1, 3, 5, 7,
	}
	for _, num := range array {
		fmt.Println(num)
	}
}

func mainP() {
	// Можно инициализировать пустой массив
	var array [3]int
	fmt.Println(array)
}
