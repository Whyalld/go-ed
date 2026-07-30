package main

import "fmt"

type age int

func (a age) isAdult() bool {
	return a >= 18
}

func main14() {
	myAge := age(20)

	fmt.Println("Пользователь совершеннолетний?", myAge.isAdult())
}
