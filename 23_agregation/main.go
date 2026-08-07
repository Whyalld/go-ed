package main

import "fmt"

type Driver struct {
	Name string
}

type SafeCar struct {
	Driver *Driver // Агрегация: Ссылка на водителя. Водитель существует отдельно
	Model  string
}

func main() {
	// Водитель существует сам по себе
	man := &Driver{Name: "Said"}

	taxi := SafeCar{
		Driver: man,
		Model:  "Toyota",
	}

	fmt.Println(taxi.Driver.Name)
}
