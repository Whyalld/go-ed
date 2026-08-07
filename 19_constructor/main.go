package main

import "fmt"

type user struct {
	name string
	age  int
}

// Прямое создание объекта (без конструктора)
var user1 = user{name: "Said", age: 18}
var user2, er = newUser("John", 25)

// Обычная функция, выполняющая роль конструктора
func newUser(name string, age int) (*user, error) {
	if age < 0 {
		return nil, fmt.Errorf("Некоректное значение возраста")
	}

	// Возвращаем указатаель на созданный объект
	return &user{name: name, age: age}, nil
}

func main18() {
	fmt.Println(user1)
	fmt.Println(*user2)
}
