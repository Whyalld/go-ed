package main

/* Классического наследования классов (когда один класс расширяет другой через extends) в Go нет.
Вместо этого Go пропагандирует принцип: «Композиция лучше наследования».

Вместо того чтобы говорить "Собака — это Животное", в Go говорят: "Собака включает в себя Животное".
Это делается с помощью встраивания полей (struct embedding). */

import "fmt"

// Базовая структура
type animal struct {
	name string
}

func (a *animal) eat() {
	fmt.Printf("%s кушает\n", a.name)
}

// Дочерняя структура
type dog struct {
	animal // Встраивание: мы просто пишем имя типа без названия поля!
	breed string
}

func (d *dog) bark() {
	fmt.Printf("%s лает: Гав-гав!\n", d.name) // Обратите внимание: d.name, а не d.animal.name!
}

func main19() {
	// Создаем объект dog
	myDog := dog{
		animal: animal{name: "Джон"}, // Заполняем встроенную структуру
		breed: "овчарка",
	}

	// Происходит "проброс" (delegation) методов и полей
	fmt.Println(myDog.breed, myDog.name) // Поле name доступно напрямую у myDog
	myDog.eat() // Метод eat доступен напрямую у myDog
	myDog.bark() // Собственный метод dog
}