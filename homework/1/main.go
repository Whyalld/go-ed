package main

import (
	"errors"
	"fmt"
	"math"
)

func homeWork() {
	per, e1 := perOfTriangle(29, 32, 12)
	area, e2 := areaOfTriangle(10, 10, 60)
	if e1 == nil && e2 == nil {
		fmt.Printf("Периметр треугольника: %v, Площадь: %v\n", per, area)
	} else {
		fmt.Println(e1, e2)
	}

}

func perOfTriangle(a uint, b uint, c uint) (uint, error) {
	if a+b < c || a+c < b || c+b < a {
		return 0, errors.New("This triangle does not exit")
	}
	return a + b + c, nil
}

func areaOfTriangle(a float64, b float64, angle float64) (float64, error) {
	if angle > 90 {
		return 0, errors.New("Angle must be bellow 90 degrees")
	}
	return 0.5 * a * b * math.Sin(math.Pi*(angle/180)), nil
}
