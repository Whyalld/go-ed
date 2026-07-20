package main

import "fmt"

// Объявляем интерфейс
type Speaker interface {
	Speak() string
}

// создаем обычную структуру
type Dog struct {
	Name string
}

// Пишем метод для структуры Dog. Сигнатура совпадает с интерфейсом!
func (d Dog) Speak() string {
	return "Гав-Гав"
}

// Функция, которая принимает ИНТЕРФЕЙС
func Introduce(s Speaker) {
	fmt.Println(s.Speak())
}

func main17() {
	myDog := Dog{Name: "Шарик"}
	
	// Передаем структуру Dog в функцию, которая ждет Speaker (интерфейс).
	// Это работает, потому что Dog автоматически имплементировал интерфейс
	Introduce(myDog)
}