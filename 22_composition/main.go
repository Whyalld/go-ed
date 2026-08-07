package main

import "fmt"

type Engine struct {
	HorsePower int
}

type Car struct {
	Engine // Композиция: Двигатель встроен "намертво" без имени поля
	Model  string
}

func main() {
	// Двигатель создается одновременно с машиной
	myCar := Car{
		Engine: Engine{HorsePower: 300},
		Model:  "Mercedes",
	}

	fmt.Println(myCar.HorsePower)
	fmt.Println(myCar.Model)
}
