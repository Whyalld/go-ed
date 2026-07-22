package main

import "fmt"

// Хочу показать в каких случаях нужно использовать интерфейс

type Fighter interface {
	shoot() string
}

func Shooting(f Fighter) {
	fmt.Println(f.shoot())
}

type SpaceShip struct{}

func (s *SpaceShip) shoot() string {
	return "Космический корабль стреляет из орудий"
}

type Man struct{}

func (m *Man) shoot() string {
	return "Человек стреляет из пистолета"
}

func main() {
	boeing := &SpaceShip{}
	julio := &Man{}

	Shooting(boeing)
	Shooting(julio)
}
