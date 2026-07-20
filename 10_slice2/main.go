package main

import "fmt"

func main6() {
	toDoList := [...]string{
		"Купить бальзам",
		"Купить банан",
		"Пойти на стрижку",
		"Купить подарок маме",
	}

	var tasks []string
	tasks = toDoList[0:4]

	for i := range tasks {
		fmt.Println(tasks[i])
	}

	// Если изменить массив, то изм-я отобразатся и на срезе взятом с массива
	toDoList[2] = "Не идти на стрижку, лучше помыть голову просто"

	println("")
	changeTasks(tasks)

	for _, value := range tasks {
		fmt.Println(value)
	}

	// Аналогично, при изменении среза, взятого с массива, массив будт менять тоже
	println("")
	for _, value := range toDoList {
		fmt.Println(value)
	}
}

func changeTasks(slice []string) {
	slice[0] = "Не покупить бальзам"
	slice[1] = "Не покупить банан"
}
