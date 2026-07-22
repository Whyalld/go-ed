package main

import (
	"fmt"
	"math"
)

type Figure interface {
	Area() int
}

func CalculateArea(f Figure) {
	fmt.Println(f.Area())
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() int {
	return int(math.Pi * c.Radius * c.Radius)
}

type Rectangle struct {
	Width  int
	Length int
}

func (r *Rectangle) Area() int {
	return r.Length * r.Width
}

func main() {
	circle := &Circle{Radius: 10}
	rectangle := &Rectangle{Length: 10, Width: 10}

	CalculateArea(circle)
	CalculateArea(rectangle)
}
