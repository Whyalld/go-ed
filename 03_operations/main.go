package main

import "fmt"

func main() {
	var pi float32 = 3.14
	circleRadius := 2 // радиус круга в сантиметрах
	// площадь круга
	circleArea := float32(circleRadius) * float32(circleRadius) * pi

	fmt.Printf("Радиус круга %d см.\n", circleRadius)
	fmt.Printf("Формула для расчета площади круга: A=πr2\n\n")

	fmt.Printf("Площадь круга: %f см. кв.", circleArea)
}
