package main

import (
	"errors"
	"fmt"
)

var Pi float32 = 3.14

func main() {
	printCircleArea(102)
}

func printCircleArea(radius int) {
	circleArea, err := calculateArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Printf("Формула для расчета площади круга: S=πr2\n")
	fmt.Printf("Площадь круга: ~%f см. кв.\n", circleArea)
}

func calculateArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус не может быть отрицательным или равным 0!")
	}
	return float32(radius) * float32(radius) * Pi, nil
}
