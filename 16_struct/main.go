package main

import "fmt"

type employee struct {
	name   string
	gender string
	age    int
	salary int
}

func main12() {
	employee1 := newEmployee("Piter", "male", 48, 3500)
	employee2 := newEmployee("Bob", "male", 19, 1500)

	employee1.changeName("Petya")

	fmt.Printf("%s\n", employee1.getInfo())
	fmt.Printf("%s", employee2.getInfo())
}

func newEmployee(name string, gender string, age int, salary int) employee {
	return employee{
		name:   name,
		gender: gender,
		age:    age,
		salary: salary,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник: %s\nПол: %s\nВозраст: %d\nЗарплата: %d\n", e.name, e.gender, e.age, e.salary)
}

func (e *employee) changeName(name string) {
	e.name = name
}
